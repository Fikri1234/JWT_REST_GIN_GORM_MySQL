package main

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/router"
	"log"

	"github.com/spf13/viper"
)

func init() {

	// read config environment
	configuration.ReadConfig()

}

func main() {

	var err error

	// Setup database
	configuration.DB, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	port := viper.GetString("PORT")

	// Setup router
	router := router.NewRoutes()
	log.Fatal(router.Run(":" + port))
}
