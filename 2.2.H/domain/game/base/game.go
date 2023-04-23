package base

import (
	"math/rand"
	"time"
)

type (
	IGameCore interface {
		Init(*Game)
		PlayGame(*Game)
		DrawCard(*Game)
		IsOver(*Game) bool
		GetWinner(*Game) IPlayer
	}

	Game struct {
		GameCore IGameCore
		Deck     IDeck
		Players  []IPlayer
		Winner   IPlayer
	}

	GameOptions struct {
		Deck     IDeck
		Players  []IPlayer
		GameCore IGameCore
	}
)

func (game *Game) Start() {
	rand.Seed(time.Now().UnixNano())
	game.Init()
	game.Deck.Shuffle()
	game.NamePlayers()
	game.DrawCard()
	game.PlayGame()
	game.SetWinner()
}

func (game *Game) SetWinner() {
	game.Winner = game.GameCore.GetWinner(game)

}

func (game *Game) Init() {
	game.GameCore.Init(game)
}

func (game *Game) NamePlayers() {
	for _, player := range game.Players {
		player.NameSelf()
	}
}

func (game *Game) PlayGame() {
	game.GameCore.PlayGame(game)
}

func (game *Game) DrawCard() {
	game.GameCore.DrawCard(game)
}

func NewGame(opts *GameOptions) *Game {
	return &Game{
		Deck:     opts.Deck,
		Players:  opts.Players,
		GameCore: opts.GameCore,
	}
}
