package shortcut

import "strings"

type Macro struct {
	commands []Command
}

func NewMacro() *Macro {
	return &Macro{
		commands: []Command{},
	}
}

func (c *Macro) Do() {
	for _, command := range c.commands {
		command.Do()
	}
}

func (c *Macro) Undo() {
	for _, command := range c.commands {
		command.Undo()
	}
}

func (c *Macro) Name() string {
	names := []string{}
	for _, command := range c.commands {
		names = append(names, command.Name())
	}
	return strings.Join(names, " & ")
}

func (c *Macro) Add(command Command) {
	c.commands = append(c.commands, command)
}
