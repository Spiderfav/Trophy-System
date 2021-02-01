package main

import (
	"fmt"

	xbl "github.com/jbowens/xbl"
)

func main() {
	client, err := xbl.Login("ruifavinha2002@gmail.com", "Infamous23")
	if err != nil {
		panic(err)
	}
	gamertag := client.Gamertag()
	fmt.Println(gamertag)
}
