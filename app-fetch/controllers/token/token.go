package token

import (
	respModel "backend-engineer-test/app-fetch/model/response"
	"backend-engineer-test/app-fetch/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	TokenRepo repository.TokenRepositoryInterface
}

func (tc *TokenController) GetClaimTokenController(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	claim, err := tc.TokenRepo.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(200, respModel.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "Generate claim token successfully",
		Data:    claim,
	})

}
