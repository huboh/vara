package app

import (
	"github.com/huboh/vara"
	"github.com/huboh/vara/pkg/modules/config"
	"github.com/huboh/vara/pkg/modules/event"
	"github.com/huboh/vara/pkg/modules/json"

	"github.com/huboh/vara/examples/rest-api/modules/auth"
	"github.com/huboh/vara/examples/rest-api/modules/database"
	"github.com/huboh/vara/examples/rest-api/modules/user"
)

type Module struct{}

func (mod *Module) Config() *vara.ModuleConfig {
	return &vara.ModuleConfig{
		Imports: []vara.Module{
			&config.Module{},
			&event.Module{},
			&database.Module{},
			&json.Module{IsGlobal: true},
			&auth.Module{},
			&user.Module{},
		},
	}
}
