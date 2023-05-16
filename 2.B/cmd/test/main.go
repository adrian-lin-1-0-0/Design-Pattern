package main

import (
	"big2/pkg/player"
	"fmt"
	"os"
)

func main() {

	player := player.NewPlayer(&player.PlayerOptions{
		Core:   player.NewHumanPlayer(),
		Writer: os.Stdout,
		Reader: os.Stdin,
	})

	player.NamePlayer()
	fmt.Println(player.Name)
}
