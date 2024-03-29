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
	}

	Game struct {
		Over     bool
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
	game.Init()
	game.Deck.Shuffle()
	game.NamePlayers()
	game.DrawCard()
	game.PlayGame()
}

func (game *Game) Init() {
	rand.Seed(time.Now().UnixNano())
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
