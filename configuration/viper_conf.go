package configuration

import (
	"github.com/spf13/viper"
)

// ReadConfig default
func ReadConfig() {
	viper.SetConfigName("properties-staging")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resource")

	viper.ReadInConfig()
}
