package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register configs file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error in reading configs: ", err)
	}

	port := viper.Get("prod.port")

	fmt.Println(port)
}
