package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	createNewFile("sample.log", "ini sample log")
	message, err := readFile("sample.log")
	if err == nil {
		fmt.Println(message)
	} else {
		fmt.Println("Error :", err.Error())
	}
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var message string

	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	return nil
}
