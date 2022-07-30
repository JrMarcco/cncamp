package framework

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type IResponse interface {
	Json(obj any) IResponse

	Redirect(path string) IResponse

	SetHeader(key string, val string) IResponse
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	SetStatus(code int) IResponse
	SetOkStatus() IResponse
}

var _ IResponse = new(Context)

func (ctx *Context) Json(obj any) IResponse {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("content-Type", "application/json")
	if _, err = ctx.rspWriter.Write(bytes); err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	return ctx
}

func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.rspWriter, ctx.req, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.rspWriter.Header().Add(key, val)
	return ctx
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.rspWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

func (ctx *Context) SetStatus(code int) IResponse {
	ctx.rspWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() IResponse {
	ctx.rspWriter.WriteHeader(http.StatusOK)
	return ctx
}
