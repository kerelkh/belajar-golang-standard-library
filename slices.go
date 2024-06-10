package main

import (
	"fmt"
	"slices"
)

func main() {
	strings := []string{"kerel", "khalif", "afif"}
	values := []int{1, 4, 2, 3, 5}

	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(strings))
	fmt.Println(slices.Min(strings))
	fmt.Println(slices.Contains(strings, "khalif"))
	fmt.Println(slices.Index(strings, "afif"))

}
