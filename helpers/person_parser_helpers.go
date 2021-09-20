package helpers

import "goExercise/models"

// parse InsertPersonRequestBody object to person
func CreatePersonObject(personRequest models.InsertPersonRequestBody) models.Person{
	
	person := models.Person{}

	person.Name = personRequest.Name
	person.HouseId = personRequest.HouseId
	person.IsMarried = personRequest.IsMarried

	return person
}