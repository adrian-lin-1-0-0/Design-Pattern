package main

import (
	"big2"
	"fmt"
)

func main() {

	player := big2.NewPlayer(&big2.PlayerOptions{
		Core: big2.NewHumanPlayer(nil),
	})

	player.NamePlayer()
	fmt.Println(player.Name)
}
