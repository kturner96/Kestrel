package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConnection string
	Port string
}

func LoadCfg() (Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load cfg.")
	}

	var cfg Config
	
	cfg.DbConnection = os.Getenv("DB_CONNECTION")
	cfg.Port = os.Getenv("PORT")

	return cfg
}