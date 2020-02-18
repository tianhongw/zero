package zero

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine is the framework's instance
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of zero.Engine
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router[getRouterKey(method, pattern)] = handler
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, found := e.router[getRouterKey(req.Method, req.URL.Path)]; found {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 not found: %s\n", req.URL)
	}
}

func getRouterKey(method, pattern string) string {
	return method + "-" + pattern
}
