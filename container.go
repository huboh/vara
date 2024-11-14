package vara

import "go.uber.org/dig"

type scope interface {
	Decorate(decorator interface{}, opts ...dig.DecorateOption) error
	Invoke(function interface{}, opts ...dig.InvokeOption) (err error)
	Provide(constructor interface{}, opts ...dig.ProvideOption) error
	Scope(name string, opts ...dig.ScopeOption) *dig.Scope
	String() string
}
