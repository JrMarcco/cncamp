package framework

type NewInstance func(...any) (any, error)

type ServiceProvider interface {
	Register(Container) NewInstance
	Boot(Container) error
	Params(Container) []any
	IsDefer() bool
	Name() string
}
