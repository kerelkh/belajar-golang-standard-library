package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `required:"true" max:"10"`
	Age  string `required:"true" max:"5"`
}

func readFile(v any) {
	vType := reflect.TypeOf(v)
	fmt.Println(vType.Name())
	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		fmt.Println("Name of field:", field.Name)
		fmt.Println("Name of type:", field.Type)
	}
}

func isValid(data any) bool {
	vType := reflect.TypeOf(data)
	result := true
	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		tag := field.Tag.Get("required")
		if tag == "true" {
			fmt.Println("Required field:", field.Name)
			vCheck := reflect.ValueOf(data).Field(i).Interface() != ""
			if !vCheck {
				return false
			}
		}
	}

	return result
}
func main() {
	person := Person{"kerel", "26"}
	readFile(person)
	sampleType := reflect.TypeOf(person)
	field := sampleType.Field(0)
	required := field.Tag.Get("required")
	fmt.Println(required)

	person2 := Person{"kere", ""}
	fmt.Println(isValid(person2))

}
