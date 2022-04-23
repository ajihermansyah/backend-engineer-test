package main

import (
	"backend-engineer-test/app-auth/config"
	userController "backend-engineer-test/app-auth/controllers/user"
	userRepository "backend-engineer-test/app-auth/repository/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConfigInit()

	router := gin.New()

	userRepo := userRepository.NewUserRepository()
	userHandler := userController.UserController{
		UserRepo: userRepo,
	}

	// endpoint for health check API
	router.GET("/ping", userHandler.PingController)

	// endpoint user API
	router.POST("/user/register", userHandler.RegisterUserController)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	router.Run(listenPort)

}
