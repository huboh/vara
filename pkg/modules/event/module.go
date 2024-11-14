package event

import "github.com/huboh/vara"

type Module struct{}

func (m *Module) Config() *vara.ModuleConfig {
	return &vara.ModuleConfig{
		IsGlobal:             true,
		ExportConstructors:   []vara.ProviderConstructor{NewService},
		ProviderConstructors: []vara.ProviderConstructor{NewConfig, NewService},
	}
}
