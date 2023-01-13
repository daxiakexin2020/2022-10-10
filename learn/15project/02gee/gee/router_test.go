package gee

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	pattern := "/api/test/:name"
	dest := parsePattern(pattern)
	fmt.Print("dest", dest) //[api test :name]
}

func TestAddRouter(t *testing.T) {
	r := newRouter()
	r.addRouter("GET", "/api/test/:name", func(c *Context) {})
	r.addRouter("GET", "/api/rule", func(c *Context) {})
	route, m := r.getRoute("GET", "/api/test/:name")
	fmt.Println(r, route.pattern, route.part, m)
}
