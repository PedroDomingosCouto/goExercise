package helpers

import "goExercise/models"

func CreateHouseObject(houseRequest models.InsertHouseRequestBody) models.House{
	
	house := models.House{}

	house.Name = houseRequest.Name
	house.Animal = houseRequest.Animal
	house.Motto = houseRequest.Motto

	return house
}