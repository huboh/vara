package vara

import (
	"context"
	"fmt"
	"sync"
)

// Lifecycle manages the lifecycle hooks registered within the application,
// enabling actions to be performed on the application's startup and graceful shutdown.
//
// Hooks can be referenced from any provider or controller in the application. The hooks
// are triggered before the application begins accepting connections and during
// graceful shutdown to release resources.
//
// Example:
//
//	func NewUserService(d *database.Service, lc *vara.Lifecycle) *UserService {
//		us := &UserService{
//			database: d,
//		}
//
//		lc.Append(vara.LifecycleHook{
//			OnStart: func(ctx context.Context) error {
//				// code to execute before the application starts accepting connections
//				return nil
//			},
//			OnStop: func(ctx context.Context) error {
//				// code to execute during application shutdown, before the application stops accepting new requests
//				return nil
//			},
//		})
//
//		return us
//	}
type Lifecycle struct {
	mutex sync.Mutex
	hooks []LifecycleHook
}

// newLifecycle creates a new Lifecycle instance.
func newLifecycle() *Lifecycle {
	return &Lifecycle{
		hooks: []LifecycleHook{},
	}
}

// Stop calls the OnStop functions of all registered hooks in reverse order.
func (l *Lifecycle) stop(ctx context.Context) error {
	var errs []error

	for i := len(l.hooks) - 1; i >= 0; i-- {
		fn := l.hooks[i].OnStop
		if fn != nil {
			err := fn(ctx)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("lifecycle stop hooks failed: %v", errs)
	}

	return nil
}

// Start calls the OnStart functions of all registered hooks.
func (l *Lifecycle) start(ctx context.Context) error {
	var errs []error

	for _, hook := range l.hooks {
		fn := hook.OnStart
		if fn != nil {
			err := fn(ctx)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("lifecycle start hooks failed: %v", errs)
	}

	return nil
}

// Append registers a new lifecycle hook.
func (l *Lifecycle) Append(hooks ...LifecycleHook) {
	l.mutex.Lock()
	l.hooks = append(l.hooks, hooks...)
	l.mutex.Unlock()
}
