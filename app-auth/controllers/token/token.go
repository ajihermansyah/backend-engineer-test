package token

import (
	httpHelper "backend-engineer-test/app-auth/helper/http"
	"backend-engineer-test/app-auth/model"
	"backend-engineer-test/app-auth/repository"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	Helper    httpHelper.HTTPHelper
	TokenRepo repository.TokenRepositoryInterface
}

func (tc *TokenController) GenerateTokenController(ctx *gin.Context) {
	var user model.User

	body, err := ctx.GetRawData()
	if err != nil {
		fmt.Println("Error read request body :", err)
		tc.Helper.SendError(ctx, 500, "Error reading request body", "error", tc.Helper.EmptyJsonMap())
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error unmarshaling json body", err)
		tc.Helper.SendError(ctx, 500, "Error unmarshaling JSON body", "error", tc.Helper.EmptyJsonMap())
		return
	}

	jwtToken, err := tc.TokenRepo.GenerateToken(user)
	if err != nil {
		tc.Helper.SendError(ctx, 500, err.Error(), "error", tc.Helper.EmptyJsonMap())
		return
	}

	dataResponse := map[string]interface{}{
		"jwt_token": jwtToken,
	}

	tc.Helper.SendSuccess(ctx, "Generate token successfully", dataResponse)
}

func (tc *TokenController) GetClaimTokenController(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	claim, err := tc.TokenRepo.ParseToken(tokenString)
	if err != nil {
		tc.Helper.SendError(ctx, 500, err.Error(), "error", tc.Helper.EmptyJsonMap())
		return
	}

	tc.Helper.SendSuccess(ctx, "Generate claim token successfully", claim)
}
