package zero

import (
	"reflect"
	"testing"
)

func buildRouter() *router {
	r := newRouter()

	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/ping", nil)
	r.addRoute("POST", "/ping", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("POST", "/add/*", nil)

	return r
}

func TestGetRoute(t *testing.T) {
	r := buildRouter()

	n, params := r.getRoute("GET", "/hello/tianhongw")

	if n.pattern != "/hello/:name" {
		t.Fail()
	}

	name, _ := params["name"]
	if name != "tianhongw" {
		t.Fail()
	}
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name/"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}
