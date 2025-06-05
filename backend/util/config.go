package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config store all configuration of the application
type Config struct {
	DBSource        string `mapstructure:"DB_SOURCE"`
	ServerAddress   string `mapstructure:"HTTP_SERVER_ADDRESS"`
	MigrationSource string `mapstructure:"MIGRATION_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	fmt.Println(path)
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
