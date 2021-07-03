package gein

import (
	"net/http"
)

// HandleFunc defines the request handler used by gein
type HandlerFunc func(c *Context)

// Engin implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// new is the constructor of gein.Engine
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engin *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engin.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engin *Engine) GET(pattern string, handler HandlerFunc) {
	engin.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engin *Engine) POST(pattern string, handler HandlerFunc) {
	engin.addRoute("POST", pattern, handler)
}

// RUN defines the method to start a http server
func (engin *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	c := newContext(res, req)
	engin.router.handle(c)
}
func (engin *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engin)
}
