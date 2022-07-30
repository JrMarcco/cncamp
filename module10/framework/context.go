package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	core           *Core
	container      Container
	providerParams []any

	req       *http.Request
	rspWriter http.ResponseWriter

	ctx          context.Context
	handlerIndex int
	handlers     HandlerChain

	timeoutFlag bool
	writerMux   *sync.Mutex

	params map[string]string
}

func NewContext(request *http.Request, responseWriter http.ResponseWriter) *Context {
	return &Context{
		req:          request,
		rspWriter:    responseWriter,
		ctx:          request.Context(),
		writerMux:    &sync.Mutex{},
		handlerIndex: -1,
	}
}

func (ctx *Context) Make(key string) (any, error) {
	return ctx.container.Make(key)
}

func (ctx *Context) MustMake(key string) any {
	return ctx.container.MustMake(key)
}

func (ctx *Context) MakeNew(key string, params []any) (any, error) {
	return ctx.container.MakeNew(key, params)
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) Request() *http.Request {
	return ctx.req
}

func (ctx *Context) ResponseWriter() http.ResponseWriter {
	return ctx.rspWriter
}

func (ctx *Context) SetTimeout() {
	ctx.timeoutFlag = true
}

func (ctx *Context) SetHandlers(handlers []HandlerFunc) {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

func (ctx *Context) Timeout() bool {
	return ctx.timeoutFlag
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.req.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Next() {
	ctx.handlerIndex++

	for ctx.handlerIndex < len(ctx.handlers) {
		ctx.handlers[ctx.handlerIndex](ctx)
		ctx.handlerIndex++
	}
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}
