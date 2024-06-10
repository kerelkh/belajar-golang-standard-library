package main

import (
	"fmt"
	"path"
)

func main() {
	// path for unix/mac, filepath for windows
	fmt.Println(path.Dir("hello/helloworld.go"))
	fmt.Println(path.Base("hello/helloworld.go"))
	fmt.Println(path.Ext("hello/helloworld.go"))
	fmt.Println(path.Join("hello", "main", "helloworld.go"))
	fmt.Println(path.Clean("hello/helloworld.go"))
}
