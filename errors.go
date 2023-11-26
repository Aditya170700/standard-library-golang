package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func main() {
	err := GetById("aditya")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Error : validation")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Error : not found")
		} else {
			fmt.Println("Error : unknown")
		}
	}
}

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "aditya" {
		return NotFoundError
	}

	return nil
}
