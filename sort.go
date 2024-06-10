package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	peoples := []Person{
		{"Alice", 42},
		{"Bob", 50},
		{"Charlie", 35},
	}

	fmt.Println(peoples)

	sort.Sort(ByAge(peoples))
	fmt.Println(peoples)

	sort.Slice(peoples, func(i, j int) bool {
		return peoples[i].Age > peoples[j].Age
	})
	fmt.Println(peoples)
}
