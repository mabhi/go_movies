package routes

import (
	"go-movies-be/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	v1 := api.Group("/v1")
	admin := v1.Group("/admin")

	v1.Get("/movie/:id", controllers.GetMovie)
	v1.Get("/movies", controllers.GetMovies)
	v1.Get("/genres", controllers.GetGenres)
	v1.Get("/movies/:genre_id", controllers.GetMoviesByGenre)
	v1.Post("/graphql", controllers.MoviesGraphql)

	admin.Post("/editmovie", controllers.InsertMovie)
	admin.Put("/editmovie/:id", controllers.EditMovie)
	admin.Delete("/deletemovie/:id", controllers.DeleteMovie)

}
