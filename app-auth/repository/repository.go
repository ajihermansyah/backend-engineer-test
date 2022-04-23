package repository

import "backend-engineer-test/app-auth/model"

type UserRepositoryInterface interface {
	CreateUser(user model.User) (psw string, err error)
	IsCheckUserExist(user model.User) (bool, error)
	GeneratePassword() (pass string, err error)
	InsertUser(user model.User) (err error)
	GetAllUsers() (users []model.User, err error)
	GetUserByPhoneAndPassword(phone, password string) (user model.User, err error)
}

type TokenRepositoryInterface interface {
	GenerateToken(user model.User) (tokenString string, err error)
	ParseToken(token string) (claim model.TokenClaims, err error)
}
