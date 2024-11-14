package vara

// ControllerConfig defines the configuration for a controller.
// E.g route patterns, guards and metadata.
type ControllerConfig struct {
	// Pattern is a root prefix appended to each route path registered
	// within the controller.
	Pattern string

	// Metadata holds arbitrary metadata associated with the controller.
	Metadata any

	// RouteConfigs lists all routes managed by the controller.
	RouteConfigs []*RouteConfig

	// Guards contains guard instances applied globally to all routes in the controller.
	Guards []Guard

	// GuardConstructors provides constructors for creating guard instances that
	// requires dependency injection.
	GuardConstructors []GuardConstructor
}
