package controllers

import (
	"goExercise/models"
	"goExercise/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetAllHouses(c *gin.Context){
	response := services.GetAllHouses()
	c.JSON(http.StatusOK, response)
}

// Controller to save stock updates to latter be sent to jetti
// returns a HTTP 200 if all objects updated
// returns a HTTP 400 if body is malformed or there is an error
func AddNewHouse(c *gin.Context){
	
	// check json body
	requestBody := new(models.House)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect or missing body",
			"messsage": err.Error(),
		})
		return
	}

	responde, status := services.SaveHouse(requestBody)

	if status == 1{
		c.JSON(http.StatusOK, responde)
		return
	}

	c.JSON(http.StatusBadRequest, responde)

}

func DeleteHouseById(c *gin.Context){
	
	houseId, err := strconv.Atoi(c.Param("houseId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "House Id is required"})
		return
	}

	response := services.DeleteHouseById(houseId)

	c.JSON(http.StatusOK, response)

}