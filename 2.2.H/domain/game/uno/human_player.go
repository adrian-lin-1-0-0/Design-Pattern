package uno

import (
	"fmt"

	"2.2.H/domain/game/base"
	"2.2.H/utils/bus"
)

type HumanPlayerCore struct {
	RequestChan chan bus.Request
	PlayerCore
}

type HumanPlayerCoreOpt struct {
	RequestChan chan bus.Request
}

func (playerCore *HumanPlayerCore) NameSelf(player base.IPlayer) {
	req := bus.Request{Cmd: "NameSelf", ReplyChan: make(chan bus.Response)}
	playerCore.RequestChan <- req
	res := <-req.ReplyChan
	player.SetName(res.Result.(string))
}

func NewHumanPlayerCore(opts *HumanPlayerCoreOpt) *HumanPlayerCore {
	return &HumanPlayerCore{
		RequestChan: opts.RequestChan,
	}
}

type HumanTakeTurnInfo struct {
	Hand    []base.ICard
	TopCard base.ICard
}

func (playerCore *HumanPlayerCore) TakeTurn(player base.IPlayer) base.ICard {
	req := bus.Request{
		Cmd:       "TakeTurn",
		ReplyChan: make(chan bus.Response),
		Data: HumanTakeTurnInfo{
			Hand:    player.ShowHand().ToSlice(),
			TopCard: playerCore.topCard,
		},
	}
	playerCore.RequestChan <- req
	res := <-req.ReplyChan
	idx := res.Result.(int)
	fmt.Println("idx : ", idx)
	if player.ShowHand().ToSlice()[idx].Compare(playerCore.topCard) == base.EQ {
		return player.ShowHand().PlayCard(idx)
	}

	return nil
}
