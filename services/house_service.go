package services

import (
	"fmt"
	"goExercise/helpers"
	"goExercise/models"
	"log"

	"github.com/gin-gonic/gin"
)

// Get All Houses
func GetAllHouses() *[]models.House{
	helpers.ConnectDatabase()
	houses := helpers.GetAllHouses()
	return houses
	
}

// Service function to save a house
// @params:
//	-requestModel: Model with request data
func SaveHouse(house *models.House) (*gin.H, int) {
	
	result := helpers.CreateHouse(*house)

	if result == -1 {
		log.Printf("Add house failed")
		return &gin.H{"Error": "Add house failed"}, -1
	}

	return &gin.H{
		"response": "House added with success.",
	}, 1
}

// Service function to delete a house
// @params:
//	-house id
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

// Service function to Get house by id and members a house
// @params:
//	-house id
func GetHouseDataAndMembers(houseId int) (*models.HousePersons, int) {
	var housePersons models.HousePersons

	//get house data by id
	houseData := helpers.GetHouseById(houseId)
	if houseData != nil {
		housePersonsData := helpers.GetPersonsByHouseId(houseId)

		housePersons.House = (*houseData)
		housePersons.Persons= (*housePersonsData)
		return &housePersons,1
	}
	
	return nil, -1
}

// service function to update a house
func UpdateOrCreateHouse(houseId int,house *models.InsertHouseRequestBody) (*gin.H,int){

	//update
	updateResult := helpers.UpdateOrCreateHouse(houseId,house)

	if updateResult == -1 {
		return &gin.H{"Error": "Update House failed"}, -1
	}else if updateResult > 1 {
		return &gin.H{"Error": fmt.Sprintf("House added with ID: %d", updateResult) }, 1
	}

	return &gin.H{
		"response": "House updated with success.",
	}, 1
}