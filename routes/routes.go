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
		api.PATCH("house/update/:houseId", controllers.UpdateHouseData)

		//Person routes
		api.GET("persons", controllers.GetAllPersons)
		api.POST("person", controllers.AddNewPerson)
		api.DELETE("person/:personId", controllers.DeletePersonById)
		api.GET("person/:personId", controllers.GetPersonData)
	}

	return r
}