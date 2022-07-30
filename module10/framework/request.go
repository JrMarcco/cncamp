package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"strconv"
)

type IRequest interface {
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) any

	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) any

	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	Form(key string) any

	BindJson(obj any) error

	BindXml(obj any) error

	GetRawData() ([]byte, error)

	Url() string
	Method() string
	Host() string
	ClientIP() string

	Headers() map[string][]string
	Header(key string) (string, bool)

	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

var _ IRequest = new(Context)

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.req != nil {
		return ctx.req.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if intValue, err := strconv.Atoi(values[0]); err == nil {
			return intValue, true
		}
	}
	return def, false
}

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if int64Value, err := strconv.ParseInt(values[0], 10, 64); err == nil {
			return int64Value, true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if float32Value, err := strconv.ParseFloat(values[0], 32); err == nil {
			return float32(float32Value), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(params) > 0 {
		if float64Value, err := strconv.ParseFloat(values[0], 64); err == nil {
			return float64Value, true
		}
	}
	return def, false
}

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(params) > 0 {
		if boolValue, err := strconv.ParseBool(values[0]); err == nil {
			return boolValue, true
		}
	}
	return def, false
}

func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

func (ctx *Context) Query(key string) any {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values[0]
	}
	return nil
}

func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if value := ctx.Param(key); value != nil {
		if intValue, ok := value.(int); ok {
			return intValue, true
		}
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if value := ctx.Param(key); value != nil {
		if int64Value, ok := value.(int64); ok {
			return int64Value, true
		}
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if value := ctx.Param(key); value != nil {
		if float32Value, ok := value.(float32); ok {
			return float32Value, true
		}
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if value := ctx.Param(key); value != nil {
		if float64Value, ok := value.(float64); ok {
			return float64Value, true
		}
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if value := ctx.Param(key); value != nil {
		if boolValue, ok := value.(bool); ok {
			return boolValue, true
		}
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if value := ctx.Param(key); value != nil {
		if stringValue, ok := value.(string); ok {
			return stringValue, true
		}
	}
	return def, false
}

func (ctx *Context) Param(key string) any {
	if ctx.params != nil {
		if value, ok := ctx.params[key]; ok {
			return value
		}
	}
	return nil
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.req != nil {
		if err := ctx.req.ParseForm(); err == nil {
			return ctx.req.PostForm
		}
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if intValue, err := strconv.Atoi(values[0]); err == nil {
			return intValue, true
		}
	}
	return def, false
}

func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if int64Value, err := strconv.ParseInt(values[0], 10, 64); err == nil {
			return int64Value, true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if float32Value, err := strconv.ParseFloat(values[0], 32); err == nil {
			return float32(float32Value), true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(params) > 0 {
		if float64Value, err := strconv.ParseFloat(values[0], 64); err == nil {
			return float64Value, true
		}
	}
	return def, false
}

func (ctx *Context) FormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(params) > 0 {
		if boolValue, err := strconv.ParseBool(values[0]); err == nil {
			return boolValue, true
		}
	}
	return def, false
}

func (ctx *Context) FormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return def, false
}

func (ctx *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

func (ctx *Context) Form(key string) any {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values[0]
	}
	return nil
}

func (ctx *Context) BindJson(obj any) error {
	if ctx.req != nil {
		// 读取 request.Body
		body, err := ioutil.ReadAll(ctx.req.Body)
		if err != nil {
			return err
		}
		// 由于 request.Body 的读取是一次性的，
		// 锁着这里需要重新填充 request.Body，否则无法再次读取数据。
		ctx.req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		if err = json.Unmarshal(body, obj); err != nil {
			return err
		}
	} else {
		return errors.New("ctx.req empty")
	}
	return nil
}

func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.req != nil {
		body, err := ioutil.ReadAll(ctx.req.Body)
		if err != nil {
			return err
		}
		ctx.req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if err = xml.Unmarshal(body, obj); err != nil {
			return err
		}
	} else {
		return errors.New("ctx.req empty")
	}
	return nil
}

func (ctx *Context) GetRawData() ([]byte, error) {
	if ctx.req != nil {
		body, err := ioutil.ReadAll(ctx.req.Body)
		if err != nil {
			return nil, err
		}
		ctx.req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil, errors.New("ctx.req empty")
}

func (ctx *Context) Url() string {
	return ctx.req.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.req.Method
}

func (ctx *Context) Host() string {
	return ctx.req.URL.Host
}

func (ctx *Context) ClientIP() string {
	ipAddress := ctx.req.Header.Get("X-Real-Ip")
	if ipAddress != "" {
		return ipAddress
	}

	ipAddress = ctx.req.Header.Get("X-Forwarded-For")
	if ipAddress != "" {
		return ipAddress
	}

	return ctx.req.RemoteAddr
}

func (ctx *Context) Headers() map[string][]string {
	return ctx.req.Header
}

func (ctx *Context) Header(key string) (string, bool) {
	if values := ctx.req.Header.Values(key); values != nil && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.req.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

func (ctx *Context) Cookie(key string) (string, bool) {
	if value, ok := ctx.Cookies()[key]; ok {
		return value, true
	}
	return "", false
}
