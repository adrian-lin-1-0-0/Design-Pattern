package main

import (
	"fmt"

	"2.2.H/domain/game/base"
	"2.2.H/domain/game/uno"
	"2.2.H/utils/bus"
)

func main() {
	requestChan := bus.NewBus()
	cmd := uno.NewCommand(requestChan)
	go cmd.Start()
	game := base.NewGame(&base.GameOptions{
		Deck:     uno.NewDeck(),
		GameCore: &uno.GameCore{},
		Players: []base.IPlayer{
			base.NewBasePlayer(uno.NewHumanPlayerCore(&uno.HumanPlayerCoreOpt{
				RequestChan: requestChan,
			})),
			base.NewBasePlayer(uno.NewAIPlayerCore()),
			base.NewBasePlayer(uno.NewAIPlayerCore()),
			base.NewBasePlayer(uno.NewAIPlayerCore())},
	})
	game.Start()
	fmt.Println("Winner : ", game.Winner.(*base.BasePlayer).Name())
}
