package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Db struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Sslmode  string `mapstructure:"sslmode"`
}

type Config struct {
	Server *Server `mapstructure:"server"`
	Db     *Db     `mapstructure:"db"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName(os.Getenv("ENV"))
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Cannot read config %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &cfg
}
