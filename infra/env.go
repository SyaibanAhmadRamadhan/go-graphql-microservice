package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed load env | err : %v", err)
	}

	ServiceOneUrl = os.Getenv("SERVICE_ONE_URL")
	ServiceTwoUrl = os.Getenv("SERVICE_TWO_URL")
	ServiceThreeUrl = os.Getenv("SERVICE_THREE_URL")

	ServiceOneKey = os.Getenv("SERVICE_ONE_KEY")
	ServiceTwoKey = os.Getenv("SERVICE_TWO_KEY")
	ServiceThreeKey = os.Getenv("SERVICE_THREE_KEY")
	Key = os.Getenv("KEY")
}

var (
	ServiceOneUrl   string
	ServiceTwoUrl   string
	ServiceThreeUrl string
)

var (
	ServiceOneKey   string
	ServiceTwoKey   string
	ServiceThreeKey string
	Key             string
)
