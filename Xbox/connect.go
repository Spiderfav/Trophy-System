package main

import (
	"fmt"

	xbl "github.com/jbowens/xbl"
)

func main() {
	/*
		client, err := xbl.Login("ruifavinha2002@gmail.com", "Infamous23")
		if err != nil {
			panic(err)
		}
	*/
	//gamertag := client.Gamertag()
	//fmt.Println(gamertag)
	client2, err := xbl.Login("Jessgood123@hotmail.co.uk", "Daddy2020?!")
	if err != nil {
		panic(err)
	}
	gamertag2 := client2.Gamertag()
	fmt.Println(gamertag2)
	jessXID := client2.UserID()

	profile, err := client2.Profile(gamertag2)
	if err != nil {
		panic(err)
	}
	fmt.Println(profile)

	achievements, err := client2.Achievements(jessXID)
	if err != nil {
		panic(err)
	}

	fmt.Println(achievements)

}
