package services

import (
	"goExercise/helpers"
	"goExercise/models"
)

func GetAllHouses() *[]models.House{
	helpers.ConnectDatabase()
	houses := helpers.GetAllHouses()

	return houses
	
}