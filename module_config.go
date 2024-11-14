package vara

// ModuleConfig provides configuration settings for a module's functionality,
// such as its providers, controllers and imported modules.
type ModuleConfig struct {
	// IsGlobal indicates whether the module's exported providers are globally
	// available to every other module without needing an explicit import.
	//
	// This is useful for shared utilities or database connections.
	IsGlobal bool

	// Imports specifies other modules that this module depends on.
	// Providers from these modules will be available within this module.
	Imports []Module

	// Exports lists providers from this module that should be accessible to
	// other modules that import this module.
	Exports []Provider

	// ExportConstructors lists constructors for providers that should be accessible
	// in other modules importing this module.
	ExportConstructors []ProviderConstructor

	// Providers lists the providers within the module that are shared across
	// the module's other providers.
	Providers []Provider

	// ProviderConstructors lists constructors for providers that the Vara injector
	// will create and share within this module.
	ProviderConstructors []ProviderConstructor

	// Controllers lists the handlers defined in this module, which handle
	// HTTP requests and define the module's endpoints.
	Controllers []Controller

	// ControllerConstructors lists constructors for controllers in this module that
	// will be instantiated by the Vara injector.
	ControllerConstructors []ControllerConstructor
}
