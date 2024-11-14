package vara

import (
	"cmp"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/dig"
)

// Controller is any type that can receive inbound requests and produce responses.
type Controller interface {
	Config() *ControllerConfig
}

// ControllerConstructor is a function type that creates Controller instances. It may have dependencies as
// parameters and returns instances of the Controller interface, optionally returning an error on failure.
//
// Any arguments that the constructor has are treated as its dependencies. The dependencies are instantiated
// in an unspecified order along with any dependencies that they might have.
type ControllerConstructor constructor

// controller is a wrapper for managing an instance of a Controller.
type controller struct {
	Controller
	module *module
	routes []*route
	guards []*guard
}

const (
	// defaultPath is the default path prefix used
	// when no specific pattern is set in the controller config.
	defaultPath = "/"

	// pathSeparator is the separator used in route paths.
	pathSeparator = "/"
)

func newController(c Controller, m *module) (*controller, error) {
	ctrl := &controller{
		module:     m,
		Controller: c,
	}

	err := ctrl._registerGuards()
	if err != nil {
		return nil, fmt.Errorf("error registering guards: %w", err)
	}

	err = ctrl._registerRoutes()
	if err != nil {
		return nil, err
	}

	return ctrl, nil
}

// getPath constructs the full path for a route
// by combining the controller's root pattern with the route's pattern.
func (c *controller) getPath(r route) string {
	root := cmp.Or(c.Config().Pattern, defaultPath)
	path := strings.TrimPrefix(r.Pattern, pathSeparator)

	return strings.TrimSpace(
		fmt.Sprintf(
			"%s %s", r.Method, strings.Join([]string{root, path}, pathSeparator),
		),
	)
}

// getGuards retrieves the list of guards for a given route,
// including both controller-scoped guards and route-scoped guards.
func (c *controller) getGuards(r route) []*guard {
	return append(c.guards, r.guards...)
}

func (c *controller) getHandler(r route) http.Handler {
	var (
		guards  = c.getGuards(r)
		handler = r.Handler
	)

	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			gCtx := newGuardCtx(*c, r, w, req)
			allowed, err := c.runGuards(gCtx, guards)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			handler.ServeHTTP(w, req)
		},
	)
}

func (c *controller) runGuards(gCtx GuardContext, guards []*guard) (bool, error) {
	for _, guard := range guards {
		allowed, err := guard.Allow(gCtx)
		if (!allowed) || (err != nil) {
			return false, err
		}
	}
	return true, nil
}

func (c *controller) _registerRoutes() error {
	return c.module.scope.Invoke(
		func(server *httpServer) error {
			for _, rCfg := range c.Config().RouteConfigs {
				// create route from config
				r, err := newRoute(rCfg, c)
				if err != nil {
					return err
				}
				c.routes = append(c.routes, r)

				// register route handler for it's path
				server.mux.Handle(c.getPath(*r), c.getHandler(*r))
			}
			return nil
		},
	)
}

func (c *controller) _registerGuards() error {
	var (
		cCfg = c.Config()
		opts = []dig.ProvideOption{
			dig.As(new(Guard)),
			dig.Group(groupGuards.String()),
		}
	)

	for _, grd := range cCfg.Guards {
		err := c.module.scope.Provide(func() Guard { return grd }, opts...)
		if err != nil {
			return fmt.Errorf("error providing controller guard (%T): %w", grd, err)
		}
	}

	for _, grdCtor := range cCfg.GuardConstructors {
		err := c.module.scope.Provide(grdCtor, opts...)
		if err != nil {
			return fmt.Errorf("error providing controller guard (%T): %w", grdCtor, err)
		}
	}

	return c.module.scope.Invoke(
		func(input guardGroupInput) error {
			for _, grd := range input.Guards {
				g, err := newGuard(grd)
				if err != nil {
					return err
				}
				c.guards = append(c.guards, g)
			}
			return nil
		},
	)
}
