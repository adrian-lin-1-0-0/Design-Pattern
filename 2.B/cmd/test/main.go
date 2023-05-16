package main

import (
	"big2/pkg/player"
	"fmt"
)

func main() {

	player := player.NewPlayer(&player.PlayerOptions{
		Core: player.NewHumanPlayer(nil),
	})

	player.NamePlayer()
	fmt.Println(player.Name)
}
