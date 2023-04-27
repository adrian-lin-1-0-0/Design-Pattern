package uno

import (
	"2.2.H/utils/bus"
)

type Handler interface {
	Execute(bus.Request)
}

func HandlerFactory() Handler {
	nameSelf := &NameSelfHandler{}
	takeTurn := &TakeTurnHandler{}
	nameSelf.Next = takeTurn
	return nameSelf
}

type Command struct {
	RequestChan chan bus.Request
	Handler     Handler
}

func NewCommand(req chan bus.Request) *Command {
	return &Command{
		RequestChan: req,
		Handler:     HandlerFactory(),
	}
}

func (cmd *Command) Start() {
	for {
		req := <-cmd.RequestChan
		cmd.Handler.Execute(req)
	}
}
