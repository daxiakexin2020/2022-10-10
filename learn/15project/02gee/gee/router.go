package gee

import (
	"fmt"
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

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	var parts []string
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern) //[ a , b ,c ]
	key := method + "-" + pattern  //a/b/c
	_, ok := r.roots[method]       //GET
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {

	fmt.Println("******************path****************", path)
	//searchParts [api test ]
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	//找到相应的节点
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		//part => api test :name
		for index, part := range parts {

			fmt.Println("******************************index******************************", index, parts, part, searchParts)
			//part中，第一个字符，例如    :name => :
			if part[0] == ':' {
				//part[1:] 把：截取掉，留下key=>name 作为params的key
				//[hello :name]   => [hello zz]   一一对对应上
				params[part[1:]] = searchParts[index]
				fmt.Println("*******************searchParts[index]**********", index, params, part[1:], searchParts[index])
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
