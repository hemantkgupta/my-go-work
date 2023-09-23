package main

import "fmt"

func modifyMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modifySlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modifyMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println(s)
}
