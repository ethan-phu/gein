package main

import (
	"gee/gein"
	"net/http"
)

func IndexFun(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("httlowlll"))
}

func main() {
	htp := gein.New()
	htp.GET("/", IndexFun)
	htp.Run(":9000")
}
