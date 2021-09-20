package services

import (
	"fmt"
	"goExercise/helpers"
	"goExercise/models"

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
func SaveHouse(house *models.InsertHouseRequestBody) *models.House {
	
	result := helpers.CreateHouse(*house)
	
	return result
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
			return &gin.H{"Error": "Delete house failed"}
		}
	
		return &gin.H{
			"message": "ok",
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
func UpdateOrCreateHouse(houseId int,house *models.InsertHouseRequestBody) (*gin.H,*models.House){

	//update
	updateResult := helpers.UpdateOrCreateHouse(houseId,house)

	if updateResult == nil {
		return &gin.H{"Error": "Update House failed"}, nil
	}

	return &gin.H{
		"response": "House updated with success.",
	}, updateResult
}