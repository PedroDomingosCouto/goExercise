package services

import (
	"goExercise/helpers"
	"goExercise/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAllHouses() *[]models.House{
	helpers.ConnectDatabase()
	houses := helpers.GetAllHouses()

	return houses
	
}

func SaveHouse(house *models.House) (*gin.H, int) {
	
	result := helpers.CreateHouse(*house)

	if result == -1 {
		log.Printf("Add house failed")
		return &gin.H{"Error": "Add house failed"}, -1
	}

	return &gin.H{
		"response": "All stock updated with success.",
	}, 1

}