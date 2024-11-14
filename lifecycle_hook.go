package vara

import "context"

// LifecycleHook represents a lifecycle hook with functions that can be registered
// to execute specific tasks during the application's lifecycle.
type LifecycleHook struct {
	// OnStop defines a lifecycle hook that is triggered during a graceful shutdown,
	// before the application stops accepting new requests.
	//
	// Use this hook to clean up resources like closing database connections,
	// flushing logs or saving state.
	OnStop LifecycleHookFunc

	// OnStart defines a lifecycle hook that is called before the application
	// begins accepting new connections.
	//
	// Use this to set up necessary resources like database connections or caches.
	OnStart LifecycleHookFunc
}

// LifecycleHookFunc represents a [lifecycleHook] function
type LifecycleHookFunc func(context.Context) error
