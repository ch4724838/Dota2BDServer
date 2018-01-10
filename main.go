package main

import (
	"github.com/Dota2BDServer/app"
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()
	api.Get("/", app.TestAPI)
	//http://localhost:8080
	api.Run(iris.Addr(":8080"))

}
