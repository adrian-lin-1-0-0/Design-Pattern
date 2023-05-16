package main

import "big2/pkg/game"

func main() {
	bigTwo := game.NewBigTwo(&game.BigTwoOptions{
		PlayerCount: 4,
	})

	bigTwo.Run()
}
