package auth

import (
	"github.com/huboh/vara"
)

type Module struct{}

func (m *Module) Config() *vara.ModuleConfig {
	return &vara.ModuleConfig{
		ProviderConstructors: []vara.ProviderConstructor{
			newService,
			newListener,
		},
		ControllerConstructors: []vara.ControllerConstructor{
			newController,
		},
	}
}
