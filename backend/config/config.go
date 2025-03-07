package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	MAX_ORDER_ITEMS               = 10
	ICED_SURCHARGE                = 0.50
	TAKEAWAY_SURCHARGE            = 0.50
	TELEGRAM_INIT_DATA_EXPIRATION = 24 * time.Hour // 1 day
	TG_AUTH_OUTPUT_KEY            = "tg-authOutput"
)

type Env struct {
	TelegramToken string `mapstructure:"TELEGRAM_TOKEN"`
	DbURI         string `mapstructure:"DB_URI"`
	Port          string `mapstructure:"PORT"`
}

func LoadEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Could not find .env", err.Error())
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Cannot load .env", err.Error())
	}

	return &env
}
