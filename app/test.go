package app

import "github.com/kataras/iris/context"

func TestAPI(c context.Context) {
	c.HTML("<h1>Hello World!</h1>")
}
