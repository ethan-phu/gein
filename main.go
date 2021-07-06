package main

import (
	"gee/gein"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gein.HandlerFunc {
	return func(c *gein.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gein.New()
	r.Use(gein.Logger()) // global midlleware
	r.GET("/", func(c *gein.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gein</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gein.Context) {
			// expect /hello/geinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
