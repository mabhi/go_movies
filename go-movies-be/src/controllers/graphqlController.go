package controllers

import (
	"encoding/json"
	"fmt"
	"go-movies-be/src/graphiQL"
	"go-movies-be/src/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

var movies []*models.Movie

func MoviesGraphql(ctx *fiber.Ctx) error {

	query := string(ctx.Body())
	log.Printf("query is %s \n", query)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphiQL.Fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return fmt.Errorf("failed to create new schema, error: %v", err)
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		return fmt.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("output is %s \n", rJSON)

	return ctx.JSON(r.Data)
}
