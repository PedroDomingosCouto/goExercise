package controllers

import (
	"goExercise/models"
	"goExercise/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPersons(c *gin.Context){
	response := services.GetAllPersons()
	c.JSON(http.StatusOK, response)
}


// Controller to add new person
// returns a HTTP 200 if person is added
// returns a HTTP 400 if body is malformed or there is an error
func AddNewPerson(c *gin.Context){
	
	// check json body
	requestBody := new(models.InsertPersonRequestBody)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect or missing body",
			"messsage": err.Error(),
		})
		return
	}else{
		
	}

	responde, status := services.AddPerson(requestBody)

	if status != nil  {
		c.JSON(http.StatusOK, status)
		return
	}
	c.JSON(http.StatusBadRequest, responde)

}

// Controller to delete a person
// returns a HTTP 200 if person is deleted
// returns a HTTP 400 if missing person Id
func DeletePersonById(c *gin.Context){
	personId, err := strconv.Atoi(c.Param("personId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person Id is required"})
		return
	}

	response := services.DeletePersonById(personId)

	c.JSON(http.StatusOK, response)
}

// Controller to get a person by id
// returns a HTTP 200 if get person
// returns a HTTP 400 if missing house Id
func GetPersonData(c *gin.Context){
	personId, err := strconv.Atoi(c.Param("personId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person Id is required"})
		return
	}
	response := services.GetPersonData(personId)
	
	if response == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This id not exist"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Controller to update a perosn by id
// returns a HTTP 200 if person married status is updated 
// returns a HTTP 400 if missing person Id
func UpdateMarriedStatus(c *gin.Context){
	personId, err := strconv.Atoi(c.Param("personId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person Id is required"})
		return
	}

	marriedStatus, errMarried := strconv.ParseBool(c.Param("marriedStatus"))
	if errMarried != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "married status is required"})
		return
	}

	response, person := services.UpdateMarriedStatus(personId,marriedStatus)

	if person == nil {
		c.JSON(http.StatusBadRequest, response )
	}

	c.JSON(http.StatusOK, person)
}

// Controller to update a perosn by id
// returns a HTTP 200 if person house status is updated 
// returns a HTTP 400 if missing person Id
func UpdateHouseID(c *gin.Context){
	personId, err := strconv.Atoi(c.Param("personId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person Id is required"})
		return
	}

	newHouseId, errNewHouseID := strconv.Atoi(c.Param("newHouseID"))
	if errNewHouseID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New house id is required"})
		return
	}

	response, person := services.UpdateHouseId(personId, newHouseId)

	if person == nil {
		c.JSON(http.StatusBadRequest, response )
	}

	c.JSON(http.StatusOK, person)
}