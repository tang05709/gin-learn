package main

import (
	"myweb/app/core"
	"myweb/app/routers"
)

func main() {
	core.Connection()
	router := routers.InitRouter()
	//静态资源
	router.Run(":8081")
}