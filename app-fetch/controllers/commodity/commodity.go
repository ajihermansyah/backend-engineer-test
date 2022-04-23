package commodity

import (
	respModel "backend-engineer-test/app-fetch/model/response"
	"backend-engineer-test/app-fetch/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CommodityController struct {
	CommodityRepo repository.CommodityRepositoryInterface
	TokenRepo     repository.TokenRepositoryInterface
}

func (c *CommodityController) GetListCommodityController(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	_, err := c.TokenRepo.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	commodities, err := c.CommodityRepo.GetListCommodity()
	if err != nil {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, respModel.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "Get list data commodity successfully",
		Data:    commodities,
	})
}

func (c *CommodityController) GetCommodityAggregateController(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	tokenClaim, err := c.TokenRepo.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	if strings.ToLower(tokenClaim.Role) != "admin" {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: "This API can access for role admin only",
		})
		return
	}
	commodities, err := c.CommodityRepo.GetCommodityAggregate()
	if err != nil {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, respModel.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "Get list data commodity aggregate successfully",
		Data:    commodities,
	})
}
