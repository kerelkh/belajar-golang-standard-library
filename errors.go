package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
	DataNullError   = errors.New("Data Not Found")
)

func getId(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "kerel" {
		return NotFoundError
	}

	return nil
}

func getUser(uuid string) (interface{}, error) {
	if uuid == "" {
		return 0, ValidationError
	} else if uuid == "1" {
		user := map[string]string{
			"name": "Kerel Khalif Afif",
		}

		return user, nil
	} else {
		return 0, DataNullError
	}
}

func main() {
	err := getId("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknonw Error")
		}
	}

	user, err := getUser("1")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, DataNullError) {
			fmt.Println("Data Not Found")
		} else {
			fmt.Println("Unknonw Error")
		}
	} else {
		fmt.Println(user)
	}
}
