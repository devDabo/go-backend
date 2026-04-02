package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *Config

func Load() *Config {
	if cfg != nil {
		return cfg
	}
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_PORT", ":8100")
	BindAllKeys()
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	cfg = &c // cache it
	return cfg
}
