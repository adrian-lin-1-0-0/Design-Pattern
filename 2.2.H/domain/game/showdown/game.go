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

		if player, ok := game.Players[FinxMaxCardIdx(showdownList)].(*base.BasePlayer).
			PlayerCore.(interface{ AddPoint() }); ok {
			player.AddPoint()
		}
	}

EndGame:
	game.Over = true
	game.Winner = gameCore.GetWinner(game)

}

func (gameCore *GameCore) GetWinner(game *base.Game) base.IPlayer {

	playerCore, ok := game.Players[0].(*base.BasePlayer).PlayerCore.(interface{ GetPoint() int })
	if !ok {
		return nil
	}

	return findWinner(game.Players[1:], game.Players[0], playerCore)
}

func findWinner(player []base.IPlayer, winner base.IPlayer, winnerCore interface{ GetPoint() int }) base.IPlayer {
	if len(player) == 0 {
		return winner
	}

	if playerCore, ok := player[0].(*base.BasePlayer).PlayerCore.(interface{ GetPoint() int }); ok {
		if playerCore.GetPoint() > winnerCore.GetPoint() {
			winner, winnerCore = player[0], playerCore
		}
	}

	return findWinner(player[1:], winner, winnerCore)
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
