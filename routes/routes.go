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
		//House routes
		api.GET("houses", controllers.GetAllHouses)
		api.POST("house", controllers.AddNewHouse)
		api.DELETE("house/:houseId", controllers.DeleteHouseById)
		api.GET("house/:houseId", controllers.GetHouseData)


		//Person routes
		api.GET("person", controllers.GetAllPersons)
		api.POST("person", controllers.AddNewPerson)
		api.DELETE("person/:personId", controllers.DeletePersonById)
	}

	return r
}