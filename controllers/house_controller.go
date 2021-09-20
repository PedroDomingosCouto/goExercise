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

// Controller to add new house
// returns a HTTP 200 if house is added
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

	// call SaveHouse to handle logic
	responde, status := services.SaveHouse(requestBody)

	if status == 1{
		c.JSON(http.StatusOK, responde)
		return
	}
	c.JSON(http.StatusBadRequest, responde)
}

// Controller to delete a house
// returns a HTTP 200 if house is deleted
// returns a HTTP 400 if missing house Id
func DeleteHouseById(c *gin.Context){
	
	houseId, err := strconv.Atoi(c.Param("houseId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "House Id is required"})
		return
	}
	// call DeleteHouseById to handle logic
	response := services.DeleteHouseById(houseId)

	c.JSON(http.StatusOK, response)
}

// Controller to get a house by id
// returns a HTTP 200 if get house
// returns a HTTP 400 if missing house Id
func GetHouseData(c *gin.Context) {
	houseId, err := strconv.Atoi(c.Param("houseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "House Id is required"})
		return
	}
	//call GetHouseDataAndMembers to handle logic
	response, status := services.GetHouseDataAndMembers(houseId)

	if status == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this identification does not exist or was not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Controller to update a house by id
// returns a HTTP 200 if house is updated or created
// returns a HTTP 400 if missing house Id
func UpdateHouseData(c *gin.Context){
	houseId, err := strconv.Atoi(c.Param("houseId"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "House Id is required in url"})
		return
	}

	// check json body
	requestBody := new(models.InsertHouseRequestBody)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect or missing body",
			"messsage": err.Error(),
		})
		return
	}
		//call UpdateOrCreateHouse to handle logic
	responde, status := services.UpdateOrCreateHouse(houseId,requestBody)

	if status == 1{
		c.JSON(http.StatusOK, responde)
		return
	}
	c.JSON(http.StatusBadRequest, responde)
	
}