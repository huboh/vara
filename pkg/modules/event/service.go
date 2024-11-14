package event

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

var (
	ErrTimeout          = errors.New("event processing timeout")
	ErrInvalidListener  = errors.New("invalid listener function")
	ErrListenerNotFound = errors.New("listener not found")
)

type Service struct {
	config    Config
	mu        sync.RWMutex
	listeners map[string][]*Listener
}

func NewService(cfg Config) *Service {
	return &Service{
		config:    cfg,
		listeners: make(map[string][]*Listener),
	}
}

// Emit emits an event to all registered listeners
func (s *Service) Emit(ctx context.Context, evt string, payload any) error {
	var (
		wg    = sync.WaitGroup{}
		errs  = make(chan error, len(s.listeners[evt]))
		event = Event{
			Payload: payload,
			Metadata: EventMetadata{
				Name:      evt,
				Context:   ctx,
				CreatedAt: time.Now(),
			},
		}
		rmvOnceLtn = func(ltn *Listener) error {
			s.mu.RUnlock()
			defer s.mu.RLock()

			if ltn.Once {
				err := s.RemoveListener(ltn)
				if (err != nil) && (!errors.Is(err, ErrListenerNotFound)) {
					return fmt.Errorf("could not remove once listener: %w", err)
				}
			}
			return nil
		}
	)

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, ltn := range s.listeners[evt] {
		if !ltn.Async {
			var err error

			err = s.processSyncListener(ltn, event)
			if err != nil {
				return err
			}

			err = rmvOnceLtn(ltn)
			if err != nil {
				return err
			}

			continue
		}

		wg.Add(1)
		go func(ltn *Listener) {
			var err error
			defer wg.Done()

			err = s.processAsyncListener(ltn, event)
			if err != nil {
				errs <- err
				return
			}

			err = rmvOnceLtn(ltn)
			if err != nil {
				errs <- err
				return
			}
		}(ltn)
	}

	// wait for all async listeners
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

// HasListeners reports if the event has at least one listener
func (s *Service) HasListeners(evt string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.listeners[evt]) >= 1
}

// AddListener registers an event listener
func (s *Service) AddListener(listeners ...*Listener) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ltn := range listeners {
		if (ltn.Func == nil) || (ltn.Event == "") {
			return ErrInvalidListener
		}

		s.listeners[ltn.Event] = append(s.listeners[ltn.Event], ltn)

		// sort by priority
		sort.Slice(
			s.listeners[ltn.Event],
			func(i, j int) bool {
				return s.listeners[ltn.Event][i].Priority > s.listeners[ltn.Event][j].Priority
			},
		)
	}

	return nil
}

// RemoveListener removes an event listener
func (s *Service) RemoveListener(ltn *Listener) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.listeners[ltn.Event]
	if !exists {
		return ErrListenerNotFound
	}

	for i, l := range s.listeners[ltn.Event] {
		if getUintPointer(l) == getUintPointer(ltn) {
			s.listeners[ltn.Event] = append(s.listeners[ltn.Event][:i], s.listeners[ltn.Event][i+1:]...)
			return nil
		}
	}

	return ErrListenerNotFound
}

// processSyncListener handles synchronous event processing
func (s *Service) processSyncListener(ltn *Listener, evt Event) error {
	return ltn.Func(evt)
}

// processAsyncListener handles async event processing
func (s *Service) processAsyncListener(ltn *Listener, evt Event) error {
	ctx := evt.Metadata.Context
	errCh := make(chan error, 1)
	timeoutCh := time.After(s.config.AsyncTimeout)

	go func() {
		errCh <- ltn.Func(evt)
	}()

	select {
	case err := <-errCh:
		return err

	case <-ctx.Done():
		return ctx.Err()

	case <-timeoutCh:
		return fmt.Errorf("%w: event %s", ErrTimeout, ltn.Event)
	}
}
