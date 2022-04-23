package token

import (
	"backend-engineer-test/app-fetch/config"
	"backend-engineer-test/app-fetch/model"
	"backend-engineer-test/app-fetch/repository"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type TokenRepository struct {
}

func NewTokenRepository() repository.TokenRepositoryInterface {
	return &TokenRepository{}
}

func (repo *TokenRepository) ParseToken(tokenString string) (tokenClaim model.TokenClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		fmt.Println("Failed parse JWT token :", err)
		err = errors.New("Failed parse JWT Token")
		return tokenClaim, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenClaim.Name = fmt.Sprintf("%v", claims["name"])
		tokenClaim.Phone = fmt.Sprintf("%v", claims["phone"])
		tokenClaim.Role = fmt.Sprintf("%v", claims["role"])
		tokenClaim.Timestamp = fmt.Sprintf("%v", claims["timestamp"])
	} else {
		fmt.Println("Failed to parse private claim token :", err)
		err = errors.New("Failed to parse private claim token")
		return tokenClaim, err
	}

	return tokenClaim, nil

}
