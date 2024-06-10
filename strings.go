package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := " test tet   "
	fmt.Println(strings.Trim(string1, " "))
	fmt.Println(strings.Contains(string1, "tes"))
	fmt.Println(strings.ToLower(string1))
	fmt.Println(strings.ToUpper(string1))
}
