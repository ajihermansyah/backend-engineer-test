package repository

import "backend-engineer-test/app-auth/model"

type UserRepositoryInterface interface {
	CreateUser(user model.User) (psw string, err error)
	IsCheckUserExist(user model.User) (bool, error)
	GeneratePassword() (pass string, err error)
}
