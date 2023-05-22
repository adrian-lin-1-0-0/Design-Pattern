package shortcut

type MainController struct {
	keyBindings map[rune]Command
	doHistory   []Command
	undoHistory []Command
}

type Command interface {
	Do()
	Undo()
}

func (c *MainController) KeyBind(key rune, command Command) {
	c.keyBindings[key] = command
}

func (c *MainController) Reset() {
	c.keyBindings = make(map[rune]Command)
}

func (c *MainController) PressKey(key rune) {
	if command, ok := c.keyBindings[key]; ok {
		command.Do()
		c.doHistory = append(c.doHistory, command)
		if len(c.undoHistory) > 0 {
			c.undoHistory = []Command{}
		}
	}
}

func (c *MainController) Undo() {
	if len(c.doHistory) > 0 {
		command := c.doHistory[len(c.doHistory)-1]
		command.Undo()
		c.doHistory = c.doHistory[:len(c.doHistory)-1]
		c.undoHistory = append(c.undoHistory, command)
	}
}

func (c *MainController) Redo() {
	if len(c.undoHistory) > 0 {
		command := c.undoHistory[len(c.undoHistory)-1]
		command.Do()
		c.undoHistory = c.undoHistory[:len(c.undoHistory)-1]
		c.doHistory = append(c.doHistory, command)
	}
}
