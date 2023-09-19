package main

import "fmt"

func main() {
	fmt.Println('\'')           // 39
	fmt.Println('\n')           // 10
	fmt.Println('\n' == 10)     // true
	fmt.Println('\n' == '\x0A') // true
	fmt.Println('\n' == '\x0a') // true
	fmt.Println('\r')           // 13
}
