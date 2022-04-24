package auth

import (
	"backend-engineer-test/app-fetch/config"
	httpHelper "backend-engineer-test/app-fetch/helper/http"
	"backend-engineer-test/app-fetch/model"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuthorizationHeader(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		err := errors.New("Missing or invalid JWT token in the request header")
		fmt.Println("Missing or invalid JWT token : ", err)

		httpHelper.SendHttpResponse(ctx, 401, "unAuthorized", err.Error(), nil)
		ctx.Abort()
		return
	}

	tokenString := strings.Split(bearerToken, " ")

	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		fmt.Println("Failed to parse JWT Token :", err)
		err = errors.New("Failed to parse JWT Token/Invalid JWT Token")

		httpHelper.SendHttpResponse(ctx, 500, "error", err.Error(), nil)
		ctx.Abort()
		return
	}

	if !token.Valid {
		httpHelper.SendHttpResponse(ctx, 401, "unAuthorized", "Invalid JWT Token", nil)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func CheckAuthorizationAdminAccess(ctx *gin.Context) {
	var tokenClaim model.TokenClaims

	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		err := errors.New("Missing or invalid JWT token in the request header")
		fmt.Println("Missing or invalid JWT token : ", err)

		httpHelper.SendHttpResponse(ctx, 401, "unAuthorized", err.Error(), nil)
		ctx.Abort()
		return
	}

	tokenString := strings.Split(bearerToken, " ")

	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		fmt.Println("Failed to parse JWT Token :", err)
		err = errors.New("Failed to parse JWT Token/Invalid JWT Token")

		httpHelper.SendHttpResponse(ctx, 500, "error", err.Error(), nil)
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenClaim.Role = fmt.Sprintf("%v", claims["role"])
	} else {
		fmt.Println("Failed to parse private claims :", err)
		err = errors.New("Failed to parse private claims")

		httpHelper.SendHttpResponse(ctx, 500, "error", err.Error(), nil)
		ctx.Abort()
		return
	}

	if strings.ToLower(tokenClaim.Role) != "admin" {
		fmt.Println("This role can't access :", err)
		message := tokenClaim.Role + " role can't access this API"
		err = errors.New(message)

		httpHelper.SendHttpResponse(ctx, 401, "unAuthorized", err.Error(), nil)
		ctx.Abort()
		return
	}

	ctx.Next()
}
