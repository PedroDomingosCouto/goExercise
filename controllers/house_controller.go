package controllers

import (
	"goExercise/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllHouses(c *gin.Context){
	response := services.GetAllHouses()
	c.JSON(http.StatusOK, response)
}