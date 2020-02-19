package zero

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	r.handlers[getRouterKey(method, pattern)] = handler
}

func (r *router) handle(c *Context) {
	if handler, found := r.handlers[getRouterKey(c.Request.Method, c.Request.URL.Path)]; found {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 not found: %s\n", c.Request.URL.Path)
	}
}

func getRouterKey(method, pattern string) string {
	return method + "-" + pattern
}
