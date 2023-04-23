package showdown

import (
	"math/rand"
	"strconv"

	"2.2.H/domain/game/base"
)

type AIPlayerCore struct {
	PlayerCore
}

func (playerCore *AIPlayerCore) NameSelf(player base.IPlayer) {
	player.SetName("AIPlayer" + strconv.Itoa(rand.Intn(100)))
}

func (playerCore *AIPlayerCore) TakeTurn(player base.IPlayer) base.ICard {
	return player.ShowHand().PlayCard(0)
}

func NewAIPlayerCore() *AIPlayerCore {
	return &AIPlayerCore{}
}
