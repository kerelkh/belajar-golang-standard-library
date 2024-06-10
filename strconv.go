package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string to int
	i, err := strconv.Atoi("-42")
	if err == nil {
		fmt.Println(i)
	}
	//int to string
	fmt.Println(strconv.Itoa(-42))

	//string to bool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	}

	//string to float
	float, err := strconv.ParseFloat("3.21", 10)
	if err == nil {
		fmt.Println(float)
	}
}
