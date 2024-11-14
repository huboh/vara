package vara

import (
	"fmt"

	"go.uber.org/dig"
)

// Module is an interface representing a self-contained unit of functionality.
// It exposes providers that can be imported by other modules within the application.
type Module interface {
	// Config returns the module config containing providers, exports
	// and imports required by the module.
	Config() *ModuleConfig
}

// module is a wrapper for managing an instance of a Module.
type module struct {
	Module
	scope       scope
	parent      *module
	imports     []*module
	controllers []*controller
}

func newModule(m Module, s scope) (*module, error) {
	var (
		err error
		mod = &module{
			scope:  s,
			Module: m,
		}
	)

	for _, imported := range m.Config().Imports {
		subMod, err := newModule(imported, mod._newChildScope(imported))
		if err != nil {
			return nil, fmt.Errorf("could not build module: %w", imported, err)
		}

		err = subMod._assignParent(mod)
		if err != nil {
			return nil, err
		}

		err = subMod._registerExportedProviders()
		if err != nil {
			return nil, fmt.Errorf("could not register exported providers: %w", err)
		}
	}

	err = mod._registerProviders()
	if err != nil {
		return nil, fmt.Errorf("could not register providers: %w", err)
	}

	err = mod._registerControllers()
	if err != nil {
		return nil, fmt.Errorf("could not register controllers: %w", err)
	}

	return mod, err
}

// _assignParent assigns the module's parent and append itself to the parent import list
func (m *module) _assignParent(parent *module) error {
	if parent != nil {
		m.parent = parent
		m.parent.imports = append(m.parent.imports, m)
	}
	return nil
}

func (m *module) _newChildScope(mod Module) scope {
	return m.scope.Scope(GetToken(mod))
}

func (m *module) _isExportedProvider(provider ProviderConstructor) bool {
	for _, export := range m.Config().ExportConstructors {
		if GetToken(export) == GetToken(provider) {
			return true
		}
	}
	return false
}

func (m *module) _registerProviders() error {
	mCfg := m.Config()
	for _, pvdCtor := range mCfg.ProviderConstructors {
		isGlobExport := (mCfg.IsGlobal && m._isExportedProvider(pvdCtor))
		// a global module's exported providers
		// should be made available to all available scopes
		err := m.scope.Provide(pvdCtor, dig.Export(isGlobExport))
		if err != nil {
			return fmt.Errorf("error providing provider (%T): %w", pvdCtor, err)
		}
	}
	return nil
}

func (m *module) _registerControllers() error {
	var (
		mCfg = m.Config()
		opts = []dig.ProvideOption{
			dig.As(new(Controller)),
			dig.Group(groupControllers.String()),
		}
	)

	for _, ctrlCtor := range mCfg.ControllerConstructors {
		err := m.scope.Provide(ctrlCtor, opts...)
		if err != nil {
			return fmt.Errorf("error providing controller (%T): %w", ctrlCtor, err)
		}
	}

	return m.scope.Invoke(
		func(input controllerGroupInput) error {
			for _, controller := range input.Controllers {
				ctrl, err := newController(controller, m)
				if err != nil {
					return err
				}
				m.controllers = append(m.controllers, ctrl)
			}
			return nil
		},
	)
}

// _registerExportedProviders registers the current module's exports in it's parent scope
func (m *module) _registerExportedProviders() error {
	mCfg := m.Config()
	// a global module's exports would be
	// available to all available scopes already.
	// so, no need to provide it to parent module's scope
	if (mCfg.IsGlobal) || (m.parent == nil) {
		return nil
	}

	for _, pvdCtor := range mCfg.ExportConstructors {
		err := m.parent.scope.Provide(pvdCtor)
		if err != nil {
			return fmt.Errorf("error providing export (%T): %w", pvdCtor, err)
		}
	}

	return nil
}
