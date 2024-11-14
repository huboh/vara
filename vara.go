package vara

import (
	"context"
	"net/http"

	"go.uber.org/dig"
)

// App represents the main application
type App struct {
	module     *module
	container  *dig.Container
	lifecycle  *Lifecycle
	httpServer *httpServer
}

// New initializes a new instance of App, configuring the root module and dependencies.
func New(module Module) (*App, error) {
	var err error

	c := dig.New()
	lc := newLifecycle()
	svr := newHttpServer(http.NewServeMux())

	err = c.Provide(func() *Lifecycle { return lc })
	if err != nil {
		return nil, err
	}

	err = c.Provide(func() *httpServer { return svr })
	if err != nil {
		return nil, err
	}

	m, err := newModule(module, c.Scope(GetToken(module)))
	if err != nil {
		return nil, err
	}

	a := &App{
		module:     m,
		container:  c,
		lifecycle:  lc,
		httpServer: svr,
	}
	err = a.httpServer.RegisterOnShutdown(a.onStop)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Listen(host, port string) error {
	err := a.onStart()
	if err != nil {
		return err
	}
	return a.httpServer.Listen(host, port)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}

func (a *App) onStop(ctx context.Context) (err error) {
	err = a.lifecycle.stop(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) onStart() (err error) {
	err = a.lifecycle.start(context.Background())
	if err != nil {
		return err
	}
	return nil
}
