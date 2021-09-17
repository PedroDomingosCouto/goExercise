package services

import (
	"fmt"
	"goExercise/helpers"
	"goExercise/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAllPersons() *[]models.Person{
	person := helpers.GetAllPersons()
	return person
	
}

func AddPerson(person *models.Person) (*gin.H, int){

	result := helpers.AddNewPerson(*person)

	if result == -1 {
		log.Printf("Add person failed")
		return &gin.H{"Error": "Add person failed"}, -1
	}

	return &gin.H{
		"response": "Person added with success.",
	}, 1
}


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
		"response": "Person deleted with success.",
	}	
	
}
