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
	requestBody := new(models.Person)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect or missing body",
			"messsage": err.Error(),
		})
		return
	}

	responde, status := services.AddPerson(requestBody)

	if status == 1{
		c.JSON(http.StatusOK, responde)
		return
	}
	c.JSON(http.StatusBadRequest, responde)

}


func DeletePersonById(c *gin.Context){
	personId, err := strconv.Atoi(c.Param("personId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person Id is required"})
		return
	}

	response := services.DeletePersonById(personId)

	c.JSON(http.StatusOK, response)
}

func GetPersonData(c *gin.Context){
	personId, err := strconv.Atoi(c.Param("personId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person Id is required"})
		return
	}

	response, status := services.GetPersonData(personId)
	if status == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Try Again"})
		return
	}

	c.JSON(http.StatusOK, response)
}