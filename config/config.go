package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int `env:"HTTP_PORT"`

	DDService string `env:"DD_SERVICE"`
	DDEnv     string `env:"DD_ENV"`

	MongoHost     string `env:"MONGO_HOST"`
	MongoUser     string `env:"MONGO_USER"`
	MongoPassword string `env:"MONGO_PASSWORD"`
	MongoName     string `env:"MONGO_NAME"`

	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresUsername string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`

	PostgresHostReplica     string `env:"POSTGRES_HOST_REPLICA"`
	PostgresPortReplica     int    `env:"POSTGRES_PORT_REPLICA"`
	PostgresUsernameReplica string `env:"POSTGRES_USER_REPLICA"`
	PostgresPasswordReplica string `env:"POSTGRES_PASSWORD_REPLICA"`
	PostgresDatabaseReplica string `env:"POSTGRES_DATABASE_REPLICA"`

	RedisHost     string `env:"REDIS_HOST"`
	RedisPort     int    `env:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

var (
	configuration Config
	mutex         sync.Once
)

func GetConfig() Config {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	var cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}

	return cfg
}
