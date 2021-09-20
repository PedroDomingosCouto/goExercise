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
		api.POST("houses", controllers.AddNewHouse)
		api.DELETE("houses/:houseId", controllers.DeleteHouseById)
		api.GET("houses/:houseId", controllers.GetHouseData)
		api.PUT("houses/update/:houseId", controllers.UpdateHouseData)

		//Person routes
		api.GET("persons", controllers.GetAllPersons)
		api.POST("persons", controllers.AddNewPerson)
		api.DELETE("persons/:personId", controllers.DeletePersonById)
		api.GET("persons/:personId", controllers.GetPersonData)
		api.PATCH("persons/:personId/:marriedStatus", controllers.UpdateMarriedStatus)
	}

	return r
}