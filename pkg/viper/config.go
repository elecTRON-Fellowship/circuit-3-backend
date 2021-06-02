package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// The values we will be reading from env files
type Config struct {
	DBDriver    string        `mapstructure:"DB_DRIVER"`
	DBSource    string        `mapstructure:"DB_Source"`
	Addr        string        `mapstructure:"ADDR"`
	JWTSecret   string        `mapstructure:"JWT_SECRET"`
	JWTDuration time.Duration `mapstructure:"JWT_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return
	}
	err = viper.Unmarshal(&config)
	return
}
