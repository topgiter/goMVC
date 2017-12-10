package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Controller("/", new(WorldController))

	app.Run(iris.Addr(":8080"))
}

type WorldController struct {
	mvc.C
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *WorldController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *WorldController) GetPing() string {
	return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *WorldController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}
