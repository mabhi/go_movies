package controllers

import (
	"context"
	"errors"
	"go-movies-be/src/database"
	"go-movies-be/src/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//Get a single movie record. parameters id of the movie
func GetMovie(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	var movie models.Movie
	movie.ID = int(id)

	theCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result := database.DB.WithContext(theCtx).First(&movie)

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "No record found",
		})
	}
	return ctx.JSON(movie)

}

//Get a single movie record. parameters id of the movie
func GetMoviesByGenre(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("genre_id"))

	var movies []models.Movie

	theCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result := database.DB.WithContext(theCtx).
		Joins("JOIN movie_genres ON movie_genres.movie_id = movies.id").
		Where("movie_genres.genre_id = ?", id).
		Find(&movies)

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "No record found",
		})
	}
	return ctx.JSON(movies)

}

//Get all movie list.
func GetMovies(ctx *fiber.Ctx) error {

	var movies []models.Movie

	theCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	database.DB.WithContext(theCtx).Find(&movies)
	return ctx.JSON(movies)
}

//Get all movie genres.
func GetGenres(ctx *fiber.Ctx) error {

	var genres []models.Genre

	theCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	database.DB.WithContext(theCtx).Find(&genres)
	return ctx.JSON(genres)
}

func InsertMovie(ctx *fiber.Ctx) error {

}

func EditMovie(ctx *fiber.Ctx) error {

}

func DeleteMovie(ctx *fiber.Ctx) error {

}
