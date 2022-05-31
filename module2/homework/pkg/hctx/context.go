package hctx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpContext struct {
	Req       *http.Request
	RspWriter http.ResponseWriter
}

func (httpCtx *HttpContext) WriteRspJson(httpStatusCode int, rsp any) error {
	httpCtx.RspWriter.WriteHeader(httpStatusCode)

	if rsp != nil {
		rspJson, err := json.Marshal(rsp)
		if err != nil {
			return err
		}

		_, err = httpCtx.RspWriter.Write(rspJson)
		return err
	}
	return nil
}

func (httpCtx *HttpContext) WriteRspStr(httpStatusCode int, rsp string) error {
	httpCtx.RspWriter.WriteHeader(httpStatusCode)

	if rsp != "" {
		_, err := httpCtx.RspWriter.Write([]byte(rsp))
		return err
	}
	return nil
}

func (httpCtx *HttpContext) WriteOkJson(rsp any) {
	if err := httpCtx.WriteRspJson(http.StatusOK, rsp); err != nil {
		httpCtx.Bad()
	}
}

func (httpCtx *HttpContext) WriteOkStr(rsp string) {
	if err := httpCtx.WriteRspStr(http.StatusOK, rsp); err != nil {
		httpCtx.Bad()
	}
}

func (httpCtx *HttpContext) Ok(rsp any) {
	if err := httpCtx.WriteRspJson(http.StatusOK, rsp); err != nil {
		httpCtx.Bad()
	}
}

func (httpCtx *HttpContext) NotFound() {
	if err := httpCtx.WriteRspStr(http.StatusNotFound, "404"); err != nil {
		httpCtx.Bad()
	}
}
func (httpCtx *HttpContext) Bad() {
	if err := httpCtx.WriteRspJson(http.StatusBadRequest, nil); err != nil {
		fmt.Printf("Fail to write rsp: %v", err)
	}
}

func (httpCtx *HttpContext) SetRspHeader(key string, value string) {
	httpCtx.RspWriter.Header().Set(key, value)
}

func (httpCtx *HttpContext) Reset(rspWriter http.ResponseWriter, req *http.Request) {
	httpCtx.RspWriter = rspWriter
	httpCtx.Req = req
}

func New() *HttpContext {
	return &HttpContext{}
}
