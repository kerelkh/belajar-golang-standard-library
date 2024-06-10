package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`k([a-z])l`)

	fmt.Println(regex.MatchString("kel"))
	fmt.Println(regex.MatchString("ael"))

	//find all string with maximum 10
	fmt.Println(regex.FindAllString("kdl kfl kgl kal kll kxl kyl", 10))
}
