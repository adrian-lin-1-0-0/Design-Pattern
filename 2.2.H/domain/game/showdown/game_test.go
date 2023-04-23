package showdown

import (
	"fmt"
	"testing"

	"2.2.H/domain/game/base"
)

func TestGameCore(t *testing.T) {
	game := base.NewGame(&base.GameOptions{
		Deck:     NewDeck(),
		GameCore: &GameCore{},
		Players: []base.IPlayer{
			base.NewBasePlayer(NewAIPlayerCore()),
			base.NewBasePlayer(NewAIPlayerCore()),
			base.NewBasePlayer(NewAIPlayerCore()),
			base.NewBasePlayer(NewAIPlayerCore())},
	})
	game.Start()
	fmt.Println(game.Winner.(*base.BasePlayer).Name())
}
