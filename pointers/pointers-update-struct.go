package main

import (
	"fmt"
	"log"
)

type Foo struct {
	Field1 string
	Field2 int
}

func main() {
	var f = Foo{
		Field1: "Hemant",
		Field2: 3,
	}
	MakeFooWrong(&f)
	fmt.Println(f.Field1)

	fa, err := MakeFooRight()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fa.Field1)
}

func MakeFooWrong(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

func MakeFooRight() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}
