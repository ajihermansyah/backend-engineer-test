package main

import (
	"backend-engineer-test/app-fetch/config"
	commodityController "backend-engineer-test/app-fetch/controllers/commodity"
	tokenController "backend-engineer-test/app-fetch/controllers/token"
	commodityRepository "backend-engineer-test/app-fetch/repository/commodity"
	tokenRepository "backend-engineer-test/app-fetch/repository/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConfigInit()

	router := gin.New()

	tokenRepo := tokenRepository.NewTokenRepository()
	tokenController := tokenController.TokenController{
		TokenRepo: tokenRepo,
	}

	commodityRepo := commodityRepository.NewCommodityRepository()
	commodityController := commodityController.CommodityController{
		CommodityRepo: commodityRepo,
		TokenRepo:     tokenRepo,
	}

	// endpoint for health check API
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is working fine...")
	})

	// endpoint auth API
	router.GET("/auth/claims-token", tokenController.GetClaimTokenController)

	// endpoint
	router.GET("/commodities", commodityController.GetListCommodityController)
	router.GET("/commodities/aggregate", commodityController.GetCommodityAggregateController)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	router.Run(listenPort)

}
