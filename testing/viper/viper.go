package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading configs: ", err)
	}

	fmt.Println("The port is: ", viper.Get("PORT"))
}
