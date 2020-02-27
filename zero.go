package zero

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

// Engine is the framework's instance
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New is the constructor of zero.Engine
func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}

	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}

	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.router.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.addRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewears...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares

	e.router.handle(c)
}
