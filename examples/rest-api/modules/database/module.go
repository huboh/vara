package database

import (
	"github.com/huboh/vara"
)

type Module struct{}

func (mod *Module) Config() *vara.ModuleConfig {
	return &vara.ModuleConfig{
		IsGlobal:               true,
		Imports:                []vara.Module{},
		ExportConstructors:     []vara.ProviderConstructor{newService},
		ProviderConstructors:   []vara.ProviderConstructor{newService},
		ControllerConstructors: []vara.ControllerConstructor{},
	}
}
