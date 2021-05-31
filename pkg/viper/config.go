package config

import (
	"log"

	"github.com/spf13/viper"
)

// The values we will be reading from env files
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_Source"`
	Addr     string `mapstructure:"ADDR"`
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
