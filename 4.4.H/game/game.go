package game

import (
	"fmt"

	"4.4.H/logger"
	"4.4.H/logger/mem"
)

type Game struct {
	logger  *logger.Logger
	players []Player
}

func New() *Game {

	return &Game{
		logger: mem.Logger.Get("app.game"),
		players: []Player{
			NewAI("AI 1"),
			NewAI("AI 2"),
			NewAI("AI 3"),
			NewAI("AI 4"),
		},
	}
}

func (g *Game) Start() {
	g.logger.Info("The game begins.")

	for _, p := range g.players {
		g.logger.Trace(fmt.Sprintf("The player *%s* begins his turn.", p.Name()))
		p.MakeDecision()
		g.logger.Trace(fmt.Sprintf("The player *%s* finishes his turn.", p.Name()))
	}

	g.logger.Debug("Game ends.")
}
