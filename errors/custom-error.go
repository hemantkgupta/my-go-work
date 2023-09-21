package main

import "fmt"

type Status1 int

const (
	InvalidLogin1 Status1 = iota + 1
	NotFound1
)

type StatusErr1 struct {
	Status  Status1
	Message string
}

func (se StatusErr1) Error() string {
	return se.Message
}

func main() {
	_, err := LoginAndGetData("", "", "abc.txt")

	if err != nil {
		fmt.Println(err)
	}
}

func LoginAndGetData(uid, pwd, file string) (interface{}, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr1{
			Status:  InvalidLogin1,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr1{
			Status:  NotFound1,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func getData(token interface{}, file string) (interface{}, interface{}) {
	return nil, ""
}

func login(uid string, pwd string) (interface{}, interface{}) {
	return "", nil
}
