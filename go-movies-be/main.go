package main

import (
	"fmt"
	"go-movies-be/src/database"
	"go-movies-be/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

const version = "1.0.0"

type config struct {
	port string
	env  string
}

type AppStatus struct {
	Version     string `json:"version"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

type Application struct {
	config
	logger *log.Logger
}

func main() {
	database.Connect()
	database.AutoMigrate()

	db, _ := database.DB.DB()

	defer db.Close()

	fiberApp := fiber.New()

	routes.Setup(fiberApp)
	fiberApp.Listen(fmt.Sprintf(":%s", os.Getenv("API_PORT")))

}
