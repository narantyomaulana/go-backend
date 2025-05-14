package routes

import (
	"github.com/gin-gonic/gin"

	authController "santrikoding/backend-api/controllers/auth"
	userController "santrikoding/backend-api/controllers/users"
	"santrikoding/backend-api/middlewares"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	// route register
	router.POST("/api/register", authController.Register)

	// route login
	router.POST("/api/login", authController.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), userController.FindUsers)

	return router
}
