package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT              string
	Key               string
	Secret            string
	RateCurrencyRatio string
	CacheExpired      int
)

func ConfigInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting environment config variables", err)
	}

	PORT = os.Getenv("SERVER_PORT")
	Key = os.Getenv("KEY")
	Secret = os.Getenv("SECRET")
	RateCurrencyRatio = os.Getenv("RATE_CURRENCY_RATIO")
	CacheExpired, _ = strconv.Atoi(os.Getenv("CACHE_EXPIRED"))

}
