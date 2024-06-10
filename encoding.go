package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//contoh encode menggunakan base64
	encoded := base64.StdEncoding.EncodeToString([]byte("Kerel Khalif Afif"))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString("CIlErhlQz780EsNv5asd2A==")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decoded))
	}
}
