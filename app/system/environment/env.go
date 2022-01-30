package environment

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ALPACA_KEY    = ""
	ALPACA_SECRET = ""
)

func Init() {
	err := godotenv.Load()

	if err != nil {
		panic("could not load env variables")
	}

	ALPACA_KEY = os.Getenv("APCA_API_KEY")
	ALPACA_SECRET = os.Getenv("APCA_API_SECRET")
}
