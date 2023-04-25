package uno

import (
	"2.2.H/domain/game/base"
)

type GameCore struct {
	Table []base.ICard
}

func (gameCore *GameCore) Init(game *base.Game) {
	//Future expansion capability
}

func (gameCore *GameCore) PlayGame(game *base.Game) {
	gameCore.Table = append(gameCore.Table, game.Deck.Draw())
	topCard := gameCore.Table[0]

	for !game.Over {
		for _, player := range game.Players {
			setTopCardToPlayer(player, topCard)

			if tmpCard := player.TakeTurn(); tmpCard != nil {
				if player.ShowHand().Len() == 0 {
					game.Winner = player
					goto EndGame
				}
				topCard = tmpCard
				gameCore.Table = append(gameCore.Table, topCard)
				continue
			}

			if game.Deck.Len() == 0 {
				game.Deck.Set(gameCore.Table[:len(gameCore.Table)-1])
				game.Deck.Shuffle()
			}
			player.AddCard(game.Deck.Draw())
		}
	}
EndGame:
	game.Over = true
}

func setTopCardToPlayer(player base.IPlayer, card base.ICard) {
	playerCore, ok := player.(*base.BasePlayer).PlayerCore.(interface{ SetTopCard(card base.ICard) })
	if !ok {
		panic("PlayerCore does not implement SetTopCard")
	}
	playerCore.SetTopCard(card)
}

func (gameCore *GameCore) DrawCard(game *base.Game) {
	handLimit := 5
	for i := 0; i < handLimit; i++ {
		for _, player := range game.Players {
			player.AddCard(game.Deck.Draw())
		}
	}
}
