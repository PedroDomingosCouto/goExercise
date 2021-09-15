package routes

import (
	"goExercise/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter is a function to initiate the router

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api") 
	{
		api.GET("houses", controllers.GetAllHouses)
		api.POST("house", controllers.AddNewHouse)
	}

	return r
}