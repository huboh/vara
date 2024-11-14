package auth

import (
	"context"

	"github.com/huboh/vara/pkg/modules/event"

	"github.com/huboh/vara/examples/rest-api/modules/database"
)

type service struct {
	events   *event.Service
	database *database.Service
}

func newService(e *event.Service, d *database.Service, l *listener) *service {
	return &service{
		events:   e,
		database: d,
	}
}

func (s *service) signin() map[string]string {
	ctx := context.Background()
	user := map[string]string{"id": "1"}

	err := s.events.Emit(ctx, eventUserSignin, user)
	if err != nil {
		// handle err
	}

	return user
}

func (s *service) signup() map[string]string {
	ctx := context.Background()
	user := map[string]string{"id": "1"}

	err := s.events.Emit(ctx, eventUserSignup, user)
	if err != nil {
		// handle err
	}

	return user
}
