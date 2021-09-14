package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter is a function to initiate the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
	}

	return r
}
