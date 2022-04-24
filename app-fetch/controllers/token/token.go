package token

import (
	httpHelper "backend-engineer-test/app-fetch/helper/http"
	"backend-engineer-test/app-fetch/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	Helper    httpHelper.HTTPHelper
	TokenRepo repository.TokenRepositoryInterface
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
