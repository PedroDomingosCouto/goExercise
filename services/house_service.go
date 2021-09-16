package services

import (
	"fmt"
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

func DeleteHouseById(houseId int) (*gin.H){

	qtdPersonsInHouse :=  helpers.GetQuantityPersonByHouseID(houseId)

	if qtdPersonsInHouse > 0{
		return &gin.H{"Error":fmt.Sprintf("You can't delete the house with id %b , because there are %d people living in it", houseId, qtdPersonsInHouse )}
	}else {
		result := helpers.DeleteHouseById(houseId)
		if result == -1 {
			log.Printf("Delete house failed")
			return &gin.H{"Error": "Delete house failed"}
		}
	
		return &gin.H{
			"response": "House deleted with success.",
		}
	}	
	
}
