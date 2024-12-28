package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	RedisAdr string
	RedisDB  int
}

func SetupEnv() (cfg AppConfig, err error) {

	godotenv.Load()

	redisAdr := os.Getenv("REDIS_ADR")
	if len(redisAdr) < 1 {
		return AppConfig{}, errors.New("env variable not found")
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return AppConfig{}, errors.New("env variable not found")
	}

	return AppConfig{RedisAdr: redisAdr, RedisDB: redisDB}, nil
}
