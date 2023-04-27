package uno

import (
	"fmt"

	"2.2.H/domain/game/base"
	"2.2.H/utils/bus"
)

type NameSelfHandler struct {
	Next Handler
}

func eqCase(next Handler, cmd string, req bus.Request, fn func(bus.Request)) {
	if req.Cmd == cmd {
		fmt.Println(cmd)
		fn(req)
	} else {
		if next != nil {
			next.Execute(req)
		}
	}
}

func (nameSelf *NameSelfHandler) Execute(req bus.Request) {
	eqCase(nameSelf.Next, "NameSelf", req, func(req bus.Request) {
		fmt.Print("Input Name :")
		req.ReplyChan <- bus.Response{Result: inputString()}
	})
}

type TakeTurnHandler struct {
	Next Handler
}

func showTopCard(topCard base.ICard) {
	unoCard := topCard.(*UnoCard)
	fmt.Println("Top Card : ", unoCard.Color.Name, unoCard.Number.Name)
}

func showHand(hand []base.ICard) {
	for i, card := range hand {
		unoCard := card.(*UnoCard)
		fmt.Println(i, ":", unoCard.Color.Name, unoCard.Number.Name)
	}
}

func (takeTurn *TakeTurnHandler) Execute(req bus.Request) {
	eqCase(takeTurn.Next, "TakeTurn", req, func(req bus.Request) {
		showTopCard(req.Data.(HumanTakeTurnInfo).TopCard)
		showHand(req.Data.(HumanTakeTurnInfo).Hand)
		fmt.Print("Input Index :")
		req.ReplyChan <- bus.Response{Result: inputInt()}
	})
}
