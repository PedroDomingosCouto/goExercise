package controllers

import (
	"goExercise/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPersons(c *gin.Context){
	response := services.GetAllPersons()
	c.JSON(http.StatusOK, response)
}