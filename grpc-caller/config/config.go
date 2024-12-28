package config

import (
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
		redisAdr = "localhost:6379"
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}

	return AppConfig{RedisAdr: redisAdr, RedisDB: redisDB}, nil
}
