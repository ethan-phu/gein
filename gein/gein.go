package gein

import (
	"fmt"
	"net/http"
)

// HandleFunc defines the request handler used by gein
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engin implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// new is the constructor of gein.Engine
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (engin *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engin.router[key] = handler
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
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engin.router[key]; ok {
		handler(res, req)
	} else {
		fmt.Fprintf(res, "404 NOT FOUND:%s\n", req.URL)
	}
}
func (engin *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engin)
}
