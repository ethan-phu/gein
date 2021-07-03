package main

import (
	"gee/gein"
)

func IndexFun(c *gein.Context) {
	c.JSON(200, gein.H{
		"code": 200,
		"msg":  "成功调度",
	})
}

func main() {
	htp := gein.New()
	htp.GET("/", IndexFun)
	htp.Run(":9000")
}
