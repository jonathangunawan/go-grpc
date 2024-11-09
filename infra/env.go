package infra

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathangunawan/go-grpc/entity"
)

func GetConfig() (*entity.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")

	return &entity.Config{
		Port: port,
	}, nil
}
