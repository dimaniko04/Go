package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	//server
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort int16  `mapstructure:"SERVER_PORT"`
	//database
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	//jwt token
	JwtSecret string `mapstructure:"JWT_SECRET"`
	//liqpay
	PublicKey  string `mapstructure:"PUBLIC_KEY"`
	PrivateKey string `mapstructure:"PRIVATE_KEY"`
}

func NewEnv(filename string, override bool) *Env {
	env := Env{}
	viper.SetConfigFile(filename)

	if override {
		viper.AutomaticEnv()
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading environment file", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error loading environment file", err)
	}

	return &env
}
