package vara

import (
	"fmt"
	"net/http"

	"go.uber.org/dig"
)

// RouteConfig defines the configuration for a route.
type RouteConfig struct {
	Method   string       // The HTTP method (e.g., GET, POST) for the route.
	Pattern  string       // The URL pattern that the route will match.
	Handler  http.Handler // The HTTP handler to process requests on this route.
	Metadata any          // Optional metadata that can be associated with the route.

	Guards            []Guard            // Guards to enforce conditions before route handling.
	GuardConstructors []GuardConstructor // Guard constructors for dynamic guard instantiation.
}

// route is a wrapper for managing route.
type route struct {
	*RouteConfig
	guards     []*guard    // Registered guards for the route.
	controller *controller // The controller that the route belongs to.
}

func newRoute(rCfg *RouteConfig, ctrl *controller) (*route, error) {
	r := &route{
		controller:  ctrl,
		RouteConfig: rCfg,
	}

	err := r._registerGuards()
	if err != nil {
		return nil, err
	}

	return r, nil
}

// _registerGuards registers all guards defined in the route configuration.
func (r *route) _registerGuards() error {
	var (
		scp  = r.controller.module.scope
		rCfg = r.RouteConfig
		opts = []dig.ProvideOption{
			dig.As(new(Guard)),
			dig.Group(groupGuards.String()),
		}
	)

	for _, grd := range rCfg.Guards {
		err := scp.Provide(func() Guard { return grd }, opts...)
		if err != nil {
			return fmt.Errorf("error providing route guard (%T): %w", grd, err)
		}
	}

	for _, grdCtor := range rCfg.GuardConstructors {
		err := scp.Provide(grdCtor, opts...)
		if err != nil {
			return fmt.Errorf("error providing route guard (%T): %w", grdCtor, err)
		}
	}

	return scp.Invoke(
		func(input guardGroupInput) error {
			for _, grd := range input.Guards {
				g, err := newGuard(grd)
				if err != nil {
					return err
				}
				r.guards = append(r.guards, g)
			}
			return nil
		},
	)
}
