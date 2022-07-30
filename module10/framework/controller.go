package framework

type HandlerFunc func(ctx *Context)

type HandlerChain []HandlerFunc
