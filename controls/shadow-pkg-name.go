package main

import "fmt"

func main() {

	fmt.Println(true)
	true := 10
	fmt.Println(true)

	x := 10
	fmt.Println(x)

	//fmt := "oops"
	// Below line will not compile
	// As there is no Println method in fmt
	//fmt.Println(fmt)

}
