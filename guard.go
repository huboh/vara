package vara

import "net/http"

// Guard is an interface that determines whether a request should be handled by
// a route handler or rejected based on specific criteria or metadata present at runtime.
type Guard interface {
	Allow(GuardContext) (bool, error)
}

// GuardContext provides the contextual information that a guard needs to make
// its decision.
//
// It encapsulates the HTTP request and response, along with
// route and controller metadata for a more informed decision-making process.
type GuardContext struct {
	// Http contains the request and response information.
	Http GuardContextHttp

	// RouteConfig contains metadata and configuration specific to the route.
	RouteConfig RouteConfig

	// ControllerConfig contains metadata and configuration for the controller.
	ControllerConfig ControllerConfig
}

// GuardContextHttp holds HTTP request and response information for GuardContext.
type GuardContextHttp struct {
	R *http.Request
	W http.ResponseWriter
}

func newGuardCtx(c controller, r route, w http.ResponseWriter, req *http.Request) GuardContext {
	return GuardContext{
		RouteConfig:      *r.RouteConfig,
		ControllerConfig: *c.Config(),
		Http: GuardContextHttp{
			R: req,
			W: w,
		},
	}
}

// GuardConstructor is a function that takes any number of dependencies
// as its parameters and returns an arbitrary number of values that meets the `Guard` interface
// and may optionally return an error to indicate that it failed to build the value.
//
// Any arguments that the constructor has are treated as its dependencies. The dependencies are instantiated
// in an unspecified order along with any dependencies that they might have.
type GuardConstructor constructor

// guard is a wrapper for managing an instance of a Guard.
type guard struct {
	Guard
}

func newGuard(g Guard) (*guard, error) {
	return &guard{
		Guard: g,
	}, nil
}
