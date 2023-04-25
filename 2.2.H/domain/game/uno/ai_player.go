package uno

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
	for idx, card := range player.ShowHand().ToSlice() {
		if card.Compare(playerCore.topCard) == base.EQ {
			return player.ShowHand().PlayCard(idx)
		}
	}
	return nil
}

func NewAIPlayerCore() *AIPlayerCore {
	return &AIPlayerCore{}
}
