package user

import (
	"backend-engineer-test/app-auth/config"
	"backend-engineer-test/app-auth/model"
	"backend-engineer-test/app-auth/repository"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sethvargo/go-password/password"
)

type UserRepository struct {
}

func NewUserRepository() repository.UserRepositoryInterface {
	return &UserRepository{}
}

func (repo *UserRepository) CreateUser(user model.User) (psw string, err error) {
	userExists, err := repo.IsCheckUserExist(user)
	if err != nil {
		fmt.Println("Failed user is already exists :", err)
		err = errors.New("Failed user is already exists")
		return "", err
	}

	if userExists {
		return "", errors.New("Username already exists, please try another name")

	} else {
		psw, err = repo.GeneratePassword()
		if err != nil {
			return "", errors.New("Failed generate password")
		}

		user.Password = psw
		err = repo.InsertUser(user)
		if err != nil {
			return "", errors.New("Failed insert new user")
		}
	}

	return psw, nil
}

func (repo *UserRepository) IsCheckUserExist(user model.User) (bool, error) {
	users, err := repo.GetAllUsers()
	if err != nil {
		fmt.Println("Failed get all data users :", err)
		err = errors.New("Failed get all users")
		return true, err
	}
	for _, val := range users {
		if val.Name == user.Name {
			return true, nil
		}
	}

	return false, nil
}

func (repo *UserRepository) GeneratePassword() (pass string, err error) {
	pass, err = password.Generate(4, 1, 0, false, false)
	if err != nil {
		fmt.Println("Error generating password :", err)
		return "", err
	}

	return pass, nil
}

func (repo *UserRepository) InsertUser(user model.User) (err error) {

	file, err := os.OpenFile(config.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed open database file :", err)
		return err
	}
	defer file.Close()

	timestamp := time.Now()
	csvWriter := csv.NewWriter(file)

	//writing data to database file csv
	csvWriter.Write([]string{user.Name, user.Phone, user.Role, user.Password, timestamp.Format("02 Jan 06 15:04")})
	csvWriter.Flush()

	return nil
}

func (repo *UserRepository) GetAllUsers() (users []model.User, err error) {

	// Opens the csv file
	file, err := os.Open(config.FileName)
	if err != nil {
		fmt.Println("Failed open database file :", err)
		err = errors.New("Failed open database file")
		return users, err
	}
	defer file.Close()

	// Read and parse the csv file into [][]string
	rows, _ := csv.NewReader(file).ReadAll()

	// Parse the result to new struct
	for _, row := range rows {
		user := model.User{
			Name:      row[0],
			Phone:     row[1],
			Role:      row[2],
			Password:  row[3],
			Timestamp: row[4],
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) GetUserByPhoneAndPassword(phone, password string) (user model.User, err error) {
	users, err := repo.GetAllUsers()
	if err != nil {
		fmt.Println("Failed get all data users :", err)
		err = errors.New("Failed to get all data users")
		return user, err
	}

	for _, val := range users {
		if val.Phone == phone && val.Password == password {
			return val, nil
		}
	}

	return user, errors.New("Failed to get user by phone and/or password")
}
