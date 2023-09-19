package main

import "fmt"

func main() {
	// The interpreted form.
	a := "Hello\nworld!\n\"你好世界\""
	fmt.Println(a)

	// The raw form.
	b := `Hello
world!
"你好世界"`
	fmt.Println(b)
}
