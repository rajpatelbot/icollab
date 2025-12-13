package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	GIN_MODE string `mapstructure:"GIN_MODE"`
	APP_PORT string `mapstructure:"APP_PORT"`

	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_SSLMODE  string `mapstructure:"DB_SSLMODE"`
}

var EnvConfig *Env

func NewEnv(fileName string) *Env {
	env := Env{}
	viper.SetConfigFile(fileName)

	// Read in environment variables that match
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading environmenet file, %s", err)
	}

	// loading env variables to struct
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Error loading environment file, %v", err)
	}

	return &env
}

func InitEnv(fileName string) {
	EnvConfig = NewEnv(fileName)
}
