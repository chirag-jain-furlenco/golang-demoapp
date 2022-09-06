package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var environmentMap = map[string]string{
	"1": "development",
	"3": "preproduction",
}

func InitializeConfig() {
	ENV, exists := os.LookupEnv("ENV")
	var ENVIndex string

	if exists == false {
		log.Println("Select Environment \n\n 1) Development/Staging 2) Production 3) Pre-Production")
		fmt.Scan(&ENVIndex)
		ENV = environmentMap[ENVIndex]
	}

	fmt.Println("ENVIRONMENT - ", ENV)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName(ENV)

	readConfigErr := viper.ReadInConfig()

	if readConfigErr != nil {
		log.Fatalln(readConfigErr)
	}
}
