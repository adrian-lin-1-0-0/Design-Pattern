package component

import (
	"big2/pkg/card/patterns"
	"big2/pkg/player"
)

type Table struct {
	TopPlay   patterns.CardPattern
	TopPlayer *player.Player
}

func NewTable() *Table {
	return &Table{}
}
