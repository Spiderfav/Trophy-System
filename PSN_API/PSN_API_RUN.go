package main

import (
	"fmt"

	psn "Trophy-System/PSN_API/API"
)

func main() {
	oauth, err := psn.Login("ruifavinha2002@icloud.com", "Infamous23")
	if err != nil {
		panic(err)
	}
	fmt.Println(oauth)
}
