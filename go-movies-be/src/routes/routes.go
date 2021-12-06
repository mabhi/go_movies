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

	admin.Post("/admin/editmovie", controllers.InsertMovie)
	admin.Put("/admin/editmovie/:id", controllers.EditMovie)
	admin.Delete("/admin/deletemovie/:id", controllers.DeleteMovie)

}
