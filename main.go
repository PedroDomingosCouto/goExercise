package main

import (
	"goExercise/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start API router
	r := routes.SetupRouter()
	r.Run()

}