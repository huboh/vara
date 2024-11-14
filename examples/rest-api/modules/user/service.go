package user

import (
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

func (user *service) getUser() map[string]string {
	return map[string]string{
		"id": "hello",
	}
}
