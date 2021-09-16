package services

import (
	"goExercise/helpers"
	"goExercise/models"
)

func GetAllPersons() *[]models.Person{
	person := helpers.GetAllPersons()
	return person
	
}