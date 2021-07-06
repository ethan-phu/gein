package main

import (
	"gee/gein"
	"net/http"
)

func main() {
	r := gein.New()
	r.GET("/index", func(c *gein.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gein.Context) {
			c.HTML(http.StatusOK, "<h1>Hello gein</h1>")
		})

		v1.GET("/hello", func(c *gein.Context) {
			// expect /hello?name=geinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gein.Context) {
			// expect /hello/geinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gein.Context) {
			c.JSON(http.StatusOK, gein.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
