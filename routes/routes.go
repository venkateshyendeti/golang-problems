package routes

import (
	"hello/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUserByID)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	return router
}
