// Package CG by chenaiquan<like.aiquan@qq.com> create on 2021/9/30 18.18
package CG

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 响应
	Response http.ResponseWriter
	Request  *http.Request
	// 路由
	Path string
	// 方法
	Method string
	// response header code
	Code int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  req,
		Path:     req.URL.Path,
		Method:   req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(status int) *Context {
	c.Code = status
	c.Response.WriteHeader(status)
	return c
}

func (c *Context) SetHeader(key string, value string) *Context {
	c.Response.Header().Set(key, value)
	return c
}

func (c *Context) String(status int, format string, values ...interface{}) *Context {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(status)
	_, _ = c.Response.Write([]byte(fmt.Sprintf(format, values...)))
	return c
}

func (c *Context) Json(status int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(status)
	encoder := json.NewEncoder(c.Response)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, _ = c.Response.Write(data)
}

func (c *Context) Html(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, _ = c.Response.Write([]byte(html))
}
