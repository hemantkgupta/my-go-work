package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := "hemantkgupta"

	se := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(se)

	sd, _ := base64.StdEncoding.DecodeString(se)
	fmt.Println(string(sd))
	fmt.Println()
}
