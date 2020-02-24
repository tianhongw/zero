package zero

import "net/http"

type RouterGroup struct {
	prefix      string
	middlewears []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine

	newGroup := &RouterGroup{
		prefix: prefix,
		parent: r,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func (r *RouterGroup) addRoute(method, pattern string, handler HandlerFunc) {
	fullPattern := r.prefix + pattern
	r.engine.router.addRoute(method, fullPattern, handler)
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodGet, pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodPost, pattern, handler)
}
