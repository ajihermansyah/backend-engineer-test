package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT   string
	Key    string
	Secret string
)

func ConfigInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting environment config variables", err)
	}

	PORT = os.Getenv("SERVER_PORT")
	Key = os.Getenv("KEY")
	Secret = os.Getenv("SECRET")

}
