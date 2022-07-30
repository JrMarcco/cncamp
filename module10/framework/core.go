package framework

import (
	"log"
	"net/http"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Core struct {
	container   Container
	router      map[string]*trieTree
	middlewares HandlerChain

	maxParams uint16
}

func NewCore() *Core {
	router := make(map[string]*trieTree, 4)
	router[GET] = NewTrieTree()
	router[POST] = NewTrieTree()
	router[PUT] = NewTrieTree()
	router[DELETE] = NewTrieTree()

	return &Core{
		container: NewServiceContainer(),
		router:    router,
	}
}

func (c *Core) Bind(sp ServiceProvider) error {
	return c.container.Bind(sp)
}

func (c *Core) IsBind(key string) bool {
	return c.container.IsBind(key)
}

func (c *Core) Get(url string, handlers ...HandlerFunc) {
	all := append(c.middlewares, handlers...)
	if err := c.router[GET].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...HandlerFunc) {
	all := append(c.middlewares, handlers...)
	if err := c.router[POST].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...HandlerFunc) {
	all := append(c.middlewares, handlers...)
	if err := c.router[PUT].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...HandlerFunc) {
	all := append(c.middlewares, handlers...)
	if err := c.router[DELETE].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Use(middlewares ...HandlerFunc) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx := NewContext(request, writer)
	routeNode := c.FindRouteNode(request)
	if routeNode == nil {
		ctx.SetStatus(http.StatusNotFound).Json("Not Found")
		return
	}

	ctx.SetParams(
		routeNode.parseParamsFromEndNode(request.URL.Path),
	)

	ctx.SetHandlers(routeNode.handlers)
	ctx.Next()
}

func (c *Core) FindRouteNode(request *http.Request) *trieTreeNode {
	reqUrl := request.URL.Path
	reqMethod := strings.ToUpper(request.Method)
	if methodRouters, ok := c.router[reqMethod]; ok {
		return methodRouters.root.findMatchNode(reqUrl)
	}
	return nil
}

func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}

func (c *Core) allocateContext() *Context {
	return &Context{
		core:           c,
		container:      c.container,
		providerParams: make([]any, 0, c.maxParams),
	}
}
