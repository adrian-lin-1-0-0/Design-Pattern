package game

import (
	"4.4.H/logger"
	"4.4.H/logger/mem"
)

type AI struct {
	name   string
	logger *logger.Logger
}

func NewAI(name string) *AI {
	return &AI{
		name:   name,
		logger: mem.Logger.Get("app.game.ai"),
	}
}

func (a *AI) MakeDecision() {
	a.logger.Trace(a.name + " starts making decisions...")
	a.logger.Warn(a.name + " decides to give up.")
	a.logger.Error("Something goes wrong when AI gives up.")
	a.logger.Trace(a.name + " completes its decision.")
}

func (a *AI) Name() string {
	return a.name
}
