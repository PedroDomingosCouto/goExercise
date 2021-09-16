package main

import (
	"goExercise/helpers"
	"goExercise/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	helpers.ConnectDatabase()
	
	// Start API router
	r := routes.SetupRouter()
	r.Run()

}