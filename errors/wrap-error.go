package main

import (
	"errors"
	"fmt"
	"os"
)

type Status3 int

const (
	InvalidLogin3 Status3 = iota + 1
	NotFound3
)

type StatusErr3 struct {
	Status  Status3
	Message string
	Err     error
}

func (se StatusErr3) Error() string {
	return se.Message
}

func (se StatusErr3) Unwrap() error {
	return se.Err
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}

func LoginAndGetData3(uid, pwd, file string) (interface{}, error) {
	token, err := login3(uid, pwd)
	if err != nil {
		return nil, StatusErr3{
			Status:  InvalidLogin3,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData3(token, file)
	if err != nil {
		return nil, StatusErr3{
			Status:  NotFound3,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}
	return data, nil
}

func getData3(token interface{}, file string) (interface{}, error) {
	return nil, nil
}

func login3(uid string, pwd string) (interface{}, error) {
	return "", nil
}
