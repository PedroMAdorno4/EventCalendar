package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Env Environment

func SetEnv() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&Env)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
