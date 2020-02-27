package zero

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

// H ...
type H map[string]interface{}

// Context ...
type Context struct {
	Writer  ResponseWriter
	Request *http.Request
	Params  map[string]string
	Path    string
	Method  string

	// middlewares
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++

	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) FormValue(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Abort() {
	c.index = math.MaxInt8
}

func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.Abort()
}
