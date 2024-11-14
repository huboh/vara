package auth

import (
	"context"
	"fmt"

	"github.com/huboh/vara"
	"github.com/huboh/vara/pkg/modules/event"

	"github.com/huboh/vara/examples/rest-api/modules/database"
)

type listener struct {
	events   *event.Service
	database *database.Service
}

func newListener(e *event.Service, d *database.Service, lc *vara.Lifecycle) *listener {
	l := &listener{
		database: d,
	}

	lc.Append(vara.LifecycleHook{
		OnStop: func(context.Context) error {
			return nil
		},
		OnStart: func(context.Context) error {
			return e.AddListener(
				&event.Listener{
					Async: true,
					Event: eventUserSignin,
					Func:  l.onUserSignin,
				},
				&event.Listener{
					Async: true,
					Event: eventUserSignup,
					Func:  l.onUserSignup,
				},
			)
		},
	})

	return l
}

func (l *listener) onUserSignup(e event.Event) error {
	fmt.Println("handling user signup event for user: %v", e.Payload)
	return nil
}

func (l *listener) onUserSignin(e event.Event) error {
	fmt.Println("handling user signin event for user: %v", e.Payload)
	return nil
}
