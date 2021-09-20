package services

import (
	"fmt"
	"goExercise/helpers"
	"goExercise/models"
	"log"

	"github.com/gin-gonic/gin"
)

// get all person
func GetAllPersons() *[]models.Person{
	person := helpers.GetAllPersons()
	return person
	
}

// Service function to save a person
// @params:
//	-requestModel: Model with request data
func AddPerson(person *models.InsertPersonRequestBody) (*gin.H, *models.Person){

	validateHouseExist := helpers.GetHouseById(person.HouseId)
	if validateHouseExist != nil {

		result := helpers.AddNewPerson(*person)

		if result == nil {
			log.Printf("Add person failed")
			return &gin.H{"Error": "Add person failed"}, nil
		}

		return &gin.H{
			"response": "Person added with success.",
		}, result
	}else{
		return &gin.H{
			"response": fmt.Sprintf("This  house ID %d not exist", person.HouseId),
		}, nil
	}
}

// Service function to delete a person
// @params:
//	-person id
func DeletePersonById(personId int) (*gin.H){

	result := helpers.DeletePersonById(personId)
	if result == -1 {
		log.Printf("Delete person failed")
		return &gin.H{"Error": "Delete Person failed"}
	}

	if result == 0 {
		return &gin.H{"Response": fmt.Sprintf("There is no such thing as a person with id %d", personId)}
	}
	
	return &gin.H{
		"message": "ok",
	}	
	
}
// Service function to Get person by id 
// @params:
//	-person id
func GetPersonData(personId int) (*models.PersonHouse){
	var personHouse models.PersonHouse
	
	personData := helpers.GetPersonById(personId)
	if personData != nil {
		houseData := helpers.GetHouseById(personData.HouseId)
		personHouse.Person= (*personData)
		personHouse.House= (*houseData)
		return &personHouse
	}
	return nil
}

// service function to update a person married Status
func UpdateMarriedStatus(personId int, marriedStatus bool) (*gin.H, *models.Person){
	personMarriedUpdated := helpers.UpdateMarriedStatus(personId, marriedStatus)

	if personMarriedUpdated == nil{
		return &gin.H{
			"response": fmt.Sprintf("The ID %d not exit", personId),
		}, nil
	}else{
		return &gin.H{
			"response": fmt.Sprintf("ID %d it has been successfully updated", personId),
		}, personMarriedUpdated
	}
}

// service function to update a person House
func UpdateHouseId(personId int, newHouseId int) (*gin.H, *models.Person){
	
	validateHouseExist := helpers.GetHouseById(personId)
	
	if validateHouseExist != nil {
	
		personUpdated := helpers.UpdateHouseId(personId, newHouseId)

		if personUpdated == nil{
			return &gin.H{
				"response": fmt.Sprintf("The ID %d not exit", personId),
			}, nil
		}else{
			return &gin.H{
				"response": fmt.Sprintf("ID %d it has been successfully updated", personId),
			}, personUpdated
		}
	}else{
		return &gin.H{
			"response": fmt.Sprintf("ID %d it has been successfully updated", personId),
		}, nil
	}
}
