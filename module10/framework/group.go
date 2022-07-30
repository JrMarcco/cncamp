package framework

type IGroup interface {
	Get(string, ...HandlerFunc)
	Post(string, ...HandlerFunc)
	Put(string, ...HandlerFunc)
	Delete(string, ...HandlerFunc)

	Use(middlewares ...HandlerFunc)

	Group(string) IGroup
}

type Group struct {
	core   *Core
	parent *Group
	prefix string

	middlewares HandlerChain
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) Group(url string) IGroup {
	childGroup := NewGroup(g.core, url)
	childGroup.parent = g
	return childGroup
}

func (g *Group) Get(url string, handlers ...HandlerFunc) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Get(url, all...)
}

func (g *Group) Post(url string, handlers ...HandlerFunc) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Post(url, all...)
}

func (g *Group) Put(url string, handlers ...HandlerFunc) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Put(url, all...)
}

func (g *Group) Delete(url string, handlers ...HandlerFunc) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Delete(url, all...)
}

func (g *Group) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) getMiddlewares() HandlerChain {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}
