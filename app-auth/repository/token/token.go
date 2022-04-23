package token

import (
	"backend-engineer-test/app-auth/config"
	"backend-engineer-test/app-auth/model"
	"backend-engineer-test/app-auth/repository"
	userRepo "backend-engineer-test/app-auth/repository/user"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type TokenRepository struct {
}

func NewTokenRepository() repository.TokenRepositoryInterface {
	return &TokenRepository{}
}

func (repo *TokenRepository) GenerateToken(user model.User) (tokenString string, err error) {
	userData, err := userRepo.NewUserRepository().GetUserByPhoneAndPassword(user.Phone, user.Password)
	if (err != nil || userData == model.User{}) {
		fmt.Println("User not found :", err)
		err = errors.New("User not found, please check your phone number and/or password")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      userData.Name,
		"phone":     userData.Phone,
		"role":      userData.Role,
		"timestamp": userData.Timestamp,
	})

	tokenString, err = token.SignedString([]byte(config.Secret))
	if err != nil {
		fmt.Println("Failed sign JWT token :", err)
		err = errors.New("Failed sign JWT token")
		return "", err
	}

	return tokenString, nil
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
