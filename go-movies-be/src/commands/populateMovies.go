package main

import (
	"fmt"
	"go-movies-be/src/common"
	"go-movies-be/src/database"
	"go-movies-be/src/models"
	"log"
	"strconv"
)

func main() {
	lines, err := common.ReadCsv("./seeders/movies.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	var dateFormat = "2006-01-02 00:00:00"
	var dateOnlyFormat = "2006-01-02"
	var movies []models.Movie
	for _, line := range lines {
		var err error
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatalf("unable to parse id : %s", line[0])
			break
		}
		year, err := strconv.Atoi(line[3])
		if err != nil {
			log.Fatalf("unable to parse year : %s", line[3])
			break
		}

		release_date, err := common.ParseTime(line[4], dateOnlyFormat)
		if err != nil {
			log.Fatalf("unable to parse release date at : %s", line[4])
			break
		}

		runtime, err := strconv.Atoi(line[5])
		if err != nil {
			log.Fatalf("unable to parse runtime : %s", line[5])
			break
		}
		rating, err := strconv.Atoi(line[6])
		if err != nil {
			log.Fatalf("unable to parse rating : %s", line[6])
			break
		}

		createAt, err := common.ParseTime(line[8], dateFormat)
		if err != nil {
			log.Fatalf("unable to parse created at : %s", line[8])
			break
		}

		updateAt, err := common.ParseTime(line[9], dateFormat)
		if err != nil {
			log.Fatalf("unable to parse updated at : %s", line[9])
			break
		}
		// id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at
		data := models.Movie{
			ID:          id,
			Title:       line[1],
			Description: line[2],
			Year:        year,
			ReleaseDate: release_date,
			Runtime:     runtime,
			Rating:      rating,
			MPAARating:  line[7],
			CreatedAt:   createAt,
			UpdatedAt:   updateAt,
		}
		movies = append(movies, data)
	}
	fmt.Println(movies)
	database.Connect()
	database.DB.Create(&movies)

}
