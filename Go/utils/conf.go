package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetConfigValue(key string) string {
	viper.SetConfigName("config.yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	var config map[string]interface{}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return config[key].(string)
}

func SetConfigValue(key string, value string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.Set(key, value)
	viper.WriteConfig()
}