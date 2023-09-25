package main

import "fmt"

func main() {
	byt := []byte("old ")
	bytes := append(byt, "falcon"...)

	fmt.Println(bytes)
	fmt.Println(string(bytes))

	copy()
	delete()
	pop()
	prepend()
	Insert()

}

func Insert() {

	a := []int{1, 2, 4, 5}

	i := 2
	e := 3

	// Use two append
	a = append(a[:i], append([]int{e}, a[i:]...)...)
	fmt.Println(a)
}

func prepend() {

	a := []int{1, 2, 3, 4, 5}
	e := 0

	// Make a slice using to be appended item
	a = append([]int{e}, a...)
	fmt.Println(a)
}

func pop() {
	a := []int{1, 2, 3, 4, 5}

	i := len(a) - 1

	// x is popped from a
	var x int
	x, a = a[i], a[:i]
	fmt.Println(x)
	fmt.Println(a)
}

func delete() {
	// Delete using append
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	i := 3
	j := 0

	// Deleting item at index i
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)

	b = append(b[:j], b[j+1:]...)
	fmt.Println(b)
}

func copy() {
	// Copy using append
	a := []int{1, 2, 3, 4}
	b := []int{}

	b = append(b, a...)

	fmt.Println(a)
	fmt.Println(b)
}
