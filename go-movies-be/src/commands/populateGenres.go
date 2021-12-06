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
	lines, err := common.ReadCsv("./seeders/genres.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	var genres []models.Genre
	var format = "2006-01-02 00:00:00"
	for _, line := range lines {
		var err error
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatalf("unable to parse ID : %s", line[0])
			break
		}
		createAt, err := common.ParseTime(line[2], format)
		if err != nil {
			log.Fatalf("unable to parse created at : %s", line[2])
			break
		}

		updateAt, err := common.ParseTime(line[3], format)
		if err != nil {
			log.Fatalf("unable to parse created at : %s", line[3])
			break
		}

		data := models.Genre{
			ID:        id,
			GenreName: line[1],
			CreatedAt: createAt,
			UpdatedAt: updateAt,
		}
		genres = append(genres, data)
	}
	fmt.Println(genres)
	database.Connect()
	database.DB.Create(&genres)

}
