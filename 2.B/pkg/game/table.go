package big2

import (
	"big2/pkg/card/patterns"
	"big2/pkg/player"
)

type Table struct {
	TopPlay   patterns.CardPattern
	TopPlayer *player.Player
}
