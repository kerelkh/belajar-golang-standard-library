package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "", "database password")
	host := flag.String("hostname", "localhost", "database hostname")
	port := flag.Int("port", 5432, "database port")

	flag.Parse()

	fmt.Println(*username, *password, *host, *port)
}
