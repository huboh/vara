package json

import "github.com/huboh/vara"

type Module struct {
	IsGlobal bool
}

func (m *Module) Config() *vara.ModuleConfig {
	return &vara.ModuleConfig{
		IsGlobal:             m.IsGlobal,
		ExportConstructors:   []vara.ProviderConstructor{NewService},
		ProviderConstructors: []vara.ProviderConstructor{NewService},
	}
}
