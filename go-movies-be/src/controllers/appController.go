package controllers

import (
	"context"
	"errors"
	"fmt"
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

	var some_movie map[string]interface{}

	// s := string(ctx.Body())
	// fmt.Printf("input is %s \n", s)

	if err := ctx.BodyParser(&some_movie); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": fmt.Sprintf("Error parsing movie body %s", err),
		})
	}

	movie, err := constructMovie(some_movie)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": fmt.Sprintf("Error constructing movie %s", err),
		})
	}

	movie.CreatedAt = time.Now()

	// fmt.Println(movie)

	database.DB.Create(&movie)

	return ctx.JSON(movie)
}

func EditMovie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	// fmt.Printf("id is %d \n", id)
	var some_movie map[string]interface{}

	search_movie := models.Movie{
		ID: int(id),
	}

	if err := ctx.BodyParser(&some_movie); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": fmt.Sprintf("Error parsing movie body %s", err),
		})
	}

	movie, err := constructMovie(some_movie)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": fmt.Sprintf("Error constructing movie %s", err),
		})
	}

	database.DB.Model(&search_movie).Updates(&movie)
	return ctx.JSON(movie)
}

func DeleteMovie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	movie := models.Movie{
		ID: int(id),
	}

	database.DB.Delete(&movie)

	return nil
}

func constructMovie(some_movie map[string]interface{}) (*models.Movie, error) {
	// --------
	val, ok := some_movie["release_date"]
	if !ok {
		return nil, fmt.Errorf("'release_date' missing from JSON")
	}

	release_date_string, ok := val.(string)
	release_date, _ := time.Parse("2006-01-02", release_date_string)
	// --------
	val, ok = some_movie["id"]
	if !ok {
		return nil, fmt.Errorf("'id' missing from JSON")
	}
	id_string, ok := val.(string)
	id, err := strconv.Atoi(id_string)
	if err != nil {
		return nil, fmt.Errorf("'id' not integer")
	}
	// --------
	val, ok = some_movie["runtime"]
	if !ok {
		return nil, fmt.Errorf("'runtime' missing from JSON")
	}
	runtime_string, ok := val.(string)
	runtime, err := strconv.Atoi(runtime_string)
	if err != nil {
		return nil, fmt.Errorf("'runtime' not integer")
	}
	// --------
	val, ok = some_movie["rating"]
	if !ok {
		return nil, fmt.Errorf("'rating' missing from JSON")
	}
	rating_string, ok := val.(string)
	rating, err := strconv.Atoi(rating_string)
	if err != nil {
		return nil, fmt.Errorf("'rating' not integer")
	}
	// --------
	val, ok = some_movie["description"]
	if !ok {
		return nil, fmt.Errorf("'description' missing from JSON")
	}
	description, ok := val.(string)
	// --------
	val, ok = some_movie["title"]
	if !ok {
		return nil, fmt.Errorf("'title' missing from JSON")
	}
	title, ok := val.(string)
	// --------
	val, ok = some_movie["mpaa_rating"]
	if !ok {
		return nil, fmt.Errorf("'mpaa_rating' missing from JSON")
	}
	mpaa_rating, ok := val.(string)
	// --------
	var movie = models.Movie{
		ID:          id,
		ReleaseDate: release_date,
		UpdatedAt:   time.Now(),
		Year:        release_date.Year(),
		Runtime:     runtime,
		Title:       title,
		Description: description,
		MPAARating:  mpaa_rating,
		Rating:      rating,
	}

	return &movie, nil
}
