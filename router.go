package zero

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	_, found := r.roots[method]
	if !found {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[getRouterKey(method, pattern)] = handler
}

func (r *router) getRoute(method, pattern string) (*node, map[string]string) {
	parts := parsePattern(pattern)
	params := make(map[string]string)

	root, found := r.roots[method]
	if !found {
		return nil, nil
	}

	n := root.search(parts, 0)
	if n == nil {
		return nil, nil
	}

	fullParts := parsePattern(n.pattern)
	for i, part := range fullParts {
		if part[0] == ':' {
			params[part[1:]] = parts[i]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(parts[i:], "/")
			break
		}
	}

	return n, params
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Request.Method, c.Request.URL.Path)
	if n != nil {
		c.Params = params
		r.handlers[getRouterKey(c.Request.Method, n.pattern)](c)
	} else {
		c.String(http.StatusNotFound, "404 not found: %s\n", c.Request.URL.Path)
	}
}

func getRouterKey(method, pattern string) string {
	return method + "-" + pattern
}

func parsePattern(pattern string) []string {
	strSlice := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range strSlice {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}
