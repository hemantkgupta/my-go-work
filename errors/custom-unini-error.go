package main

import "fmt"

type Status2 int

const (
	InvalidLogin2 Status2 = iota + 1
	NotFound2
)

type StatusErr2 struct {
	Status  Status2
	Message string
}

func (se StatusErr2) Error() string {
	return se.Message
}

func GenerateErrorBroken(flag bool) error {
	var genErr StatusErr2

	if flag {
		genErr = StatusErr2{
			Status: NotFound2,
		}
	}
	return genErr
}

func main() {
	err := GenerateErrorBroken(true)
	fmt.Println("GenerateErrorBroken(true) returns non-nil error:", err != nil)
	err = GenerateErrorBroken(false)
	fmt.Println("GenerateErrorBroken(false) returns non-nil error:", err != nil)
	err = GenerateErrorOKReturnNil(false)
	fmt.Println("GenerateErrorOKReturnNil(false) returns non-nil error:", err != nil)
	err = GenerateErrorUseErrorVar(true)
	fmt.Println("GenerateErrorUseErrorVar(false) returns non-nil error:", err != nil)
}

func GenerateErrorOKReturnNil(flag bool) error {
	if flag {
		return StatusErr2{
			Status: NotFound2,
		}
	}
	return nil
}

func GenerateErrorUseErrorVar(flag bool) error {
	// Note the genErr of type error not StatusErr
	var genErr error
	if flag {
		genErr = StatusErr2{
			Status: NotFound2,
		}
	}
	return genErr
}
