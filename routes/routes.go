package routes

import (
	"github.com/gin-contrib/cors" // Import the CORS middleware
	"github.com/gin-gonic/gin" // Import the Gin framework

	authController "santrikoding/backend-api/controllers/auth"
	userController "santrikoding/backend-api/controllers/users"
	"santrikoding/backend-api/middlewares"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// route register
	router.POST("/api/register", authController.Register)

	// route login
	router.POST("/api/login", authController.Login)

	router.GET("/api/users", middlewares.AuthMiddleware(), userController.FindUsers)

	// route create user
	router.POST("/api/users", middlewares.AuthMiddleware(), userController.CreateUser)

	// Route Find User
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), userController.FindUserById)

	// Route Update User
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), userController.UpdateUser)

	// Route Delete User
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), userController.DeleteUser)

	return router
}
