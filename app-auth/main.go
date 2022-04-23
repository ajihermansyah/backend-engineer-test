package main

import (
	"backend-engineer-test/app-auth/config"
	tokenController "backend-engineer-test/app-auth/controllers/token"
	userController "backend-engineer-test/app-auth/controllers/user"
	tokenRepository "backend-engineer-test/app-auth/repository/token"
	userRepository "backend-engineer-test/app-auth/repository/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConfigInit()

	router := gin.New()

	userRepo := userRepository.NewUserRepository()
	userController := userController.UserController{
		UserRepo: userRepo,
	}

	tokenRepo := tokenRepository.NewTokenRepository()
	tokenController := tokenController.TokenController{
		TokenRepo: tokenRepo,
	}

	// endpoint for health check API
	router.GET("/ping", userController.PingController)

	// endpoint auth API
	router.POST("/auth/register", userController.RegisterUserController)
	router.GET("/auth/generate-token", tokenController.GenerateTokenController)
	router.GET("/auth/claims-token", tokenController.GetClaimTokenController)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	router.Run(listenPort)

}
