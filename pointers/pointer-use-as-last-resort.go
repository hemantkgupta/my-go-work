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

	err := MakeFooWrong(&f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Field1)

	fr, err := MakeFooRight()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fr.Field1)
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
