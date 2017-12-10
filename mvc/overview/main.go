// file: main.go

package main

import (
	"github.com/kataras/iris/_examples/mvc/overview/datasource"
	"github.com/kataras/iris/_examples/mvc/overview/repositories"
	"github.com/kataras/iris/_examples/mvc/overview/services"
	"github.com/kataras/iris/_examples/mvc/overview/web/controllers"
	"github.com/kataras/iris/_examples/mvc/overview/web/middleware"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// Register our controllers.
	app.Controller("/hello", new(controllers.HelloController))

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie controller.
	movieService := services.NewMovieService(repo)

	app.Controller("/movies", new(controllers.MovieController),
		movieService,
		middleware.BasicAuth)

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
}
