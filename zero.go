package zero

import (
	"net/http"
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
	c := newContext(w, req)
	e.router.handle(c)
}
