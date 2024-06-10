package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	//server time
	fmt.Println(now)

	//local time
	fmt.Println(now.Local())

	//wib date
	loc, _ := time.LoadLocation("Asia/Jakarta")
	waktu := time.Date(2024, time.May, 24, 12, 2, 4, 0, loc)
	fmt.Println(waktu)
	fmt.Println(waktu.Local())

	//format golang time untuk dijadikan reference waktu
	formatter := "2006-01-02 15:05:05"

	value := "2024-05-05 15:40:22"
	valuetime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valuetime)
	}

	fmt.Println(valuetime.Month())
}
