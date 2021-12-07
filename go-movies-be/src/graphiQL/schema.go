package graphiQL

import (
	"context"
	"go-movies-be/src/database"
	"go-movies-be/src/models"
	"log"
	"strings"
	"time"

	"github.com/graphql-go/graphql"
)

var movies []*models.Movie

// graphql schema definition
var Fields = graphql.Fields{
	"movie": &graphql.Field{
		Type:        movieType,
		Description: "Get movie by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {

				var movie models.Movie
				movie.ID = int(id)

				theCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				result := database.DB.WithContext(theCtx).First(&movie)

				if result.Error != nil {
					return nil, result.Error
				}
				if result.RowsAffected > 0 {
					return movie, nil
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(movieType),
		Description: "Get all movies",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var movies []models.Movie

			theCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			database.DB.WithContext(theCtx).Find(&movies)
			return movies, nil
		},
	},
	"search": &graphql.Field{
		Type:        graphql.NewList(movieType),
		Description: "Search movies by title",
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var theList []*models.Movie
			search, ok := params.Args["titleContains"].(string)
			if ok {
				//Todo: query db
				for _, currentMovie := range movies {
					if strings.Contains(currentMovie.Title, search) {
						log.Println("Found one")
						theList = append(theList, currentMovie)
					}
				}
			}
			return theList, nil
		},
	},
}

var movieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"release_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"runtime": &graphql.Field{
				Type: graphql.Int,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
			"mpaa_rating": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
