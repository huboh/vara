# Vara [![GoDoc](https://pkg.go.dev/badge/github.com/huboh/vara)](https://pkg.go.dev/github.com/huboh/vara) [![Build Status](https://github.com/huboh/vara/actions/workflows/go.yml/badge.svg)](https://github.com/huboh/vara/actions/workflows/go.yml)

`Vara` is a powerful dependency injection framework for building modular, maintainable and testable Go applications.

## Features

- 🚀 Zero-configuration dependency injection
- 🔒 Type-safe dependency management
- 📦 Modular architecture with clear boundaries
- 🌐 Built-in HTTP routing and middleware
- ⚡ Lifecycle management for application components
- 🛡️ Flexible guard system for request authorization

## Core Concepts

Vara is built around four fundamental concepts:

- Modules: Organizational units that encapsulate related functionality
- Dependencies: Services and components managed through dependency injection
- Controllers: HTTP request handlers with built-in support for guards, filters, interceptors and middlewares
- Lifecycle Events: Hooks for managing application and component lifecycles

See the [documentation](https://pkg.go.dev/github.com/huboh/vara#section-documentation) to get started and learn more.

## Installation

Install Vara in your application with the following command.

```go
go get github.com/huboh/vara@latest
```

## Documentation

- [Getting Started](https://pkg.go.dev/github.com/huboh/vara#section-documentation)
- [Best Practices](https://pkg.go.dev/github.com/huboh/vara#hdr-Best_Practices)

## Quick example

Here's a minimal example to get you started. View the full example [here](https://github.com/huboh/vara/blob/main/examples/rest-api/modules/main.go)

```go
const (
    port = "5000"
    host = "localhost"
)

func main() {
    app, err := vara.New(&app.Module{})
    if err != nil {
        log.Fatal("failed to create app:", err)
    }

    err := app.Listen(host, port)
    if err != nil {
        log.Fatal("failed to start server:", err)
    }
}
```

## Authors

- Knowledge - [Website](https://huboh.vercel.app)

## License

Vara is released under the [MIT License](https://github.com/huboh/vara/blob/main/LICENCE).

## Contributing

Contributions are welcomed!
