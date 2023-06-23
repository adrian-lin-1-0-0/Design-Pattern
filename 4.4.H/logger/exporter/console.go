package exporter

import "os"

type Console struct {
	*Base
}

func NewConsole() *Console {
	c := &Console{&Base{}}
	c.SetOutput(os.Stdout)
	return c
}
