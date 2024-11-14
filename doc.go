// Package vara is a powerful dependency injection framework for Go that
// enables building modular, maintainable and testable applications.
//
// # Overview
//
// Vara simplifies the creation of highly maintainable Go applications through a modular architecture
// powered by dependency injection. It provides a robust foundation for building large-scale applications
// where components are loosely coupled, easily testable and highly reusable.
//
// # Core Concepts
//
// Vara is built around four fundamental concepts:
//  1. Modules - Organizational units that encapsulate related functionality
//  2. Dependencies - Services and components managed through dependency injection
//  3. Controllers - Http request handlers with built-in support for guards, filters, interceptors and middlewares
//  4. Lifecycle Events - Hooks for managing application and component lifecycles
//
// # Key Features
//
//   - Zero-configuration dependency injection
//   - Type-safe dependency management
//   - Modular architecture with clear boundaries
//   - Built-in HTTP routing and middleware
//   - Lifecycle management for application components
//   - Flexible guard system for request authorization
//
// # Getting Started
//
// To create a new Vara application:
//
//	func main() {
//		app, err := vara.New(&app.Module{})
//		if err != nil {
//			log.Fatal("failed to create app:", err)
//		}
//
//		err := app.Listen("localhost", "8080")
//		if err != nil {
//			log.Fatal("failed to start server:", err)
//		}
//	}
//
// # Dependency Injection
//
// Vara supports two types of dependency injection:
//
//  1. Constructor-based Injection
//  2. Direct Injection
//
// # Constructor-based Injection:
//
// Constructors are the building blocks of dependency injection in Vara. They are plain Go functions that:
// Accepts zero or more dependencies as parameters and return one or more values of any type and may
// optionally return an error as the last return value.
//
//   - Preferred for complex dependencies.
//   - Allows initialization logic and error handling.
//   - Supports dependency chains.
//
// Example:
//
//	func NewUserService(cache *cache.Service, database *database.Service) (*UserService, error) {
//		return &UserService{
//			cache: 	  cache,
//			database: database,
//		}, nil
//	}
//
// Any arguments that the constructor has are treated as its dependencies. The dependencies are instantiated
// in an unspecified order along with any dependencies that they might have, creating a dependency graph at runtime.
//
// Important notes about constructors:
//   - Constructors are lazy - they are only called when their return type is required somewhere in the application
//   - If a constructor is added for side effects (like registering handlers), its return type must be referenced
//     somewhere in the module for the constructor to be called
//   - Dependencies are instantiated in an unspecified order along with their own dependencies
//
// # Direct Injection:
//
// If a dependency itself does not require any other dependencies, you can opt to inject it directly without a constructor.
//
//   - Suitable for simple dependencies
//   - No constructor needed
//   - Faster initialization
//
// # Module System
//
// Modules are the building blocks of a Vara application. Each module must implement the [Module] interface:
//
//	type UserModule struct{}
//
//	func (m *UserModule) Config() *vara.ModuleConfig {
//		return &vara.ModuleConfig{
//			IsGlobal:         false,
//			Imports:          []vara.Module{&config.Module{}},
//			ExportsCtor:      []vara.ProviderConstructor{NewAuthService},
//			ProviderConstructors:   []vara.ProviderConstructor{NewAuthService},
//			ControllerConstructors: []vara.ControllerConstructor{},
//		}
//	}
//
// Module Configuration Options:
//   - [IsGlobal]:		  Makes this module's exports available to all other modules
//   - [Imports]:  		  Other modules required by this module
//   - [ExportsCtor]: 	  Subset of providers that will be available to other modules
//   - [ProviderConstructors]:   Internal services used within the module
//   - [ControllerConstructors]: HTTP controllers
//
// # Lifecycle Management
//
// Vara provides hooks for managing component lifecycles:
//
//	func NewUserService(d *database.Service, lc *vara.Lifecycle) *UserService {
//		svc := &UserService{}
//
//		lc.Append(vara.LifecycleHook{
//			OnStart: func(ctx context.Context) error {
//				// code to execute before the application starts accepting connections
//				return nil
//			},
//			OnStop: func(ctx context.Context) error {
//				// code to execute during application shutdown, before the application stops accepting new requests
//				return nil
//			},
//		})
//
//		return svc
//	}
//
// Lifecycle Events:
//   - OnStart: Called before the application starts accepting connections
//   - OnStop: Called during graceful shutdown
//
// # Controllers
//
// Controllers handles request routing and processing. They provide a structured way to define
// endpoints and their associated handlers.
//
// A controller must implement the [Controller] interface:
//
//	type UserController struct {
//		service *UserService
//	}
//
//	func (c *UserController) Config() *vara.ControllerConfig {
//		return &vara.ControllerConfig{
//			Pattern: 	"/api/v1",
//			GuardConstructors: []vara.GuardConstructor{newAuthGuard},
//			RouteConfigs: []*vara.RouteConfig{
//				{
//					Pattern: "/users",
//					Method:  http.MethodGet,
//					Handler: http.HandlerFunc(c.listUsers),
//					GuardConstructors: []vara.GuardConstructor{newRateLimitGuard},
//				},
//			},
//		}
//	}
//
// # Request Guards
//
// Guards are used to control access to controllers or individual routes,
// providing an additional layer of security by enforcing runtime or compile-time rules through the incoming request
// or controller/route metadata before handlers are executed.
//
// A guard must implement the [Guard] interface:
//
//	type AuthGuard struct {
//	    auth *AuthService
//	}
//
//	func newAuthGuard(a *AuthService) *AuthGuard {
//	    return &AuthGuard{
//	        auth: a,
//	    }
//	}
//
//	func (g *AuthGuard) Allow(gCtx vara.GuardContext) (bool, error) {
//	    validated, err := g.auth.Validate(gCtx.Http.R.Header.Get("Authorization"))
//	    if err != nil {
//	        return false, err
//	    }
//
//	    return validated, nil
//	}
//
// Guard Scopes:
//   - Route-level: Applied to specific routes only
//   - Controller-level: Applied to all routes in a controller
//
// # Complete application structure:
//
//	api/
//	├── main.go           	// Application entry point
//	├── modules/
//	│   ├── app/
//	│   │   └── module.go     // Root module
//	│   ├── auth/
//	│   │   ├── module.go     // Auth module
//	│   │   ├── service.go    // Auth service
//	│   │   └── controller.go // Auth controller
//	│   └── user/
//	│       ├── module.go     // User module
//	│       ├── service.go    // User service
//	│       └── repository.go // User repository
//	└── go.mod
//
// # Best Practices
//
// 1. Module Organization:
//   - Keep modules focused on a single responsibility
//   - Use clear naming conventions
//   - Group related functionality together
//
// 2. Dependency Management:
//   - Prefer constructor injection for complex dependencies
//   - Validate dependencies in constructors
//   - Use interfaces for better testability
//
// 3. Error Handling:
//   - Return meaningful errors from constructors
//   - Implement proper cleanup in OnStop handlers
//
// 4. Testing:
//   - Mock dependencies using interfaces
//   - Test modules in isolation
//   - Use table-driven tests
package vara // import "github.com/huboh/vara"
