package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
		message += string(line)
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func addMessageFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("sample.log", "This is sample message")
	message, err := readFile("sample.log")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}

	addMessageFile("sample.log", "\nadd sample baru")
}
