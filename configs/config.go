// This package is responsible for loading and managing configuration,
// including environment variables and database connection strings.
// You could use a package like viper to manage the configuration:
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_URL string
	Port   string
}

var Cfg Config

func LoadConfig() {
	viper.SetDefault("DB_URL", "postgres://user:password@localhost:5432/mydb")
	viper.SetDefault("PORT", "8080")

	Cfg.DB_URL = viper.GetString("DB_URL")
	Cfg.Port = viper.GetString("PORT")

	log.Printf("Config loaded: %+v\n", Cfg)
}
