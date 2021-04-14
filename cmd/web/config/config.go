package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT = 0
	DSN  = ""
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 9000
	}

	DSN = fmt.Sprintf("%s:%s@/%s?parseTime=true&charset=utf8&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
}
