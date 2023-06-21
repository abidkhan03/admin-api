package shared

import (
	"github.com/joho/godotenv"
)

func LoadConfig(filename string) error {
	return godotenv.Load(filename)
}
