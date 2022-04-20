package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
	AppPort    string `envconfig:"APP_PORT"`
	JWTSecret  string `envconfig:"JWT_SECRET"`
}

var instance Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		godotenv.Load()
		err := envconfig.Process("", &instance)
		if err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
