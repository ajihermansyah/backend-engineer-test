package token

import (
	"backend-engineer-test/app-auth/model"
	respModel "backend-engineer-test/app-auth/model/response"
	"backend-engineer-test/app-auth/repository"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	TokenRepo repository.TokenRepositoryInterface
}

func (tc *TokenController) GenerateTokenController(ctx *gin.Context) {
	var user model.User

	body, err := ctx.GetRawData()
	if err != nil {
		fmt.Println("Error read request body :", err)
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: "Error reading request body",
		})

		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error unmarshaling json body", err)
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: "Error unmarshaling JSON body",
		})

		return
	}

	jwtToken, err := tc.TokenRepo.GenerateToken(user)
	if err != nil {
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})

		return
	}

	dataResponse := map[string]interface{}{
		"jwt_token": jwtToken,
	}

	ctx.JSON(200, respModel.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "Generate token successfully",
		Data:    dataResponse,
	})
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
