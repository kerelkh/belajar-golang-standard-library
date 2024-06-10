package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := 100 * time.Second
	fmt.Println(duration1)
	duration2 := 10 * time.Millisecond
	fmt.Println(duration2)
	fmt.Println(duration1 + duration2)
}
