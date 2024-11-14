package vara

// Provider is a marker interface for types that can be provided as dependencies.
type Provider interface{}

// ProviderConstructor is a constructor function type for providers. It takes any
// required dependencies as parameters and returns instances that implement the `Provider`
// interface. Optionally, the constructor may return an error to signal that it failed
// to create the provider instance.
//
// Dependencies are resolved and injected into the constructor in an unspecified order.
// This allows for flexible and lazy-loading dependency injection where each dependency
// may itself have nested dependencies.
//
// Example:
//
//	func NewProvider(dep1 Dependency1, dep2 Dependency2) (Provider, error) {
//	    // Create and return a provider instance
//	}
//
// Any dependencies needed by the constructor will be resolved and instantiated
// by the module's DI scope.
type ProviderConstructor constructor
