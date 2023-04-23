package showdown

import (
	"2.2.H/domain/game/base"
)

type GameCore struct {
}

func (gameCore *GameCore) Init(game *base.Game) {
	//Future expansion capability
}

func (gameCore *GameCore) PlayGame(game *base.Game) {

	for !game.Over {
		showdownList := make([]base.ICard, 0)
		for _, player := range game.Players {
			if player.ShowHand().Len() == 0 {
				goto EndGame
			}
			showdownList = append(showdownList, player.TakeTurn())
		}
		idx := FinxMaxCardIdx(showdownList)

		if player, ok := game.Players[idx].(*base.BasePlayer).
			PlayerCore.(interface{ AddPoint() }); ok {
			player.AddPoint()
		}

	}

EndGame:
	game.Over = true
	game.Winner = gameCore.GetWinner(game)

}

func (gameCore *GameCore) GetWinner(game *base.Game) base.IPlayer {
	var winner base.IPlayer
	for _, IPlayer := range game.Players {
		if winner == nil {
			winner = IPlayer
		}

		if player, ok := IPlayer.(*base.BasePlayer).PlayerCore.(interface{ GetPoint() int }); ok {
			if winnerCore := winner.(*base.BasePlayer).PlayerCore.(interface{ GetPoint() int }); ok {
				if player.GetPoint() > winnerCore.GetPoint() {
					winner = IPlayer
				}
			}
		}

	}
	return winner
}

func (gameCore *GameCore) DrawCard(game *base.Game) {
	for _, player := range game.Players {
		if game.Deck.Len() == 0 {
			return
		}
		card := game.Deck.Draw()
		player.AddCard(card)
	}
	gameCore.DrawCard(game)
}
