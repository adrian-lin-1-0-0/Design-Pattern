package shortcut

import "fmt"

type Terminal struct {
	MainController    *MainController
	ConnectTelecom    ConnectTelecom
	DisconnectTelecom DisconnectTelecom
	MoveTankBackward  MoveTankBackward
	MoveTankForward   MoveTankForward
}

func NewTerminal() *Terminal {
	telecom := &Telecom{}
	tank := &Tank{}
	return &Terminal{
		MainController: &MainController{
			keyBindings: make(map[rune]Command),
		},
		ConnectTelecom: ConnectTelecom{
			Telecom: telecom,
		},
		DisconnectTelecom: DisconnectTelecom{
			Telecom: telecom,
		},
		MoveTankBackward: MoveTankBackward{
			Tank: tank,
		},
		MoveTankForward: MoveTankForward{
			Tank: tank,
		},
	}
}

func (c *Terminal) SetShortcut() {
	var macro, key rune
	fmt.Print("設置巨集指令 (y/n): ")
	fmt.Scanf("%c\n", &macro)
	fmt.Print("Key: ")
	fmt.Scanf("%c\n", &key)
	fmt.Printf("要將哪一道指令設置到快捷鍵 %c 上: \n", key)
	fmt.Println("(0) MoveTankForward\n(1) MoveTankBackward\n(2) ConnectTelecom\n(3) DisconnectTelecom\n(4) ResetMainControlKeyboard")
	var command rune
	fmt.Scanf("%c\n", &command)
	switch command {
	case '0':
		c.MainController.KeyBind(key, c.MoveTankForward)
		fmt.Printf("%c : %s \n", key, "MoveTankForward")
	case '1':
		c.MainController.KeyBind(key, c.MoveTankBackward)
		fmt.Printf("%c : %s \n", key, "MoveTankBackward")
	case '2':
		c.MainController.KeyBind(key, c.ConnectTelecom)
		fmt.Printf("%c : %s \n", key, "ConnectTelecom")
	case '3':
		c.MainController.KeyBind(key, c.DisconnectTelecom)
		fmt.Printf("%c : %s \n", key, "DisconnectTelecom")
	case '4':
		c.MainController.Reset()
		fmt.Printf("%c : %s \n", key, "ResetMainControlKeyboard")
	}
}

func (c *Terminal) Run() {
	var key rune
	for {
		fmt.Print("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵: ")
		fmt.Scanf("%c\n", &key)
		switch key {
		case '1':
			c.SetShortcut()
		case '2':
			c.MainController.Undo()
		case '3':
			c.MainController.Redo()
		default:
			c.MainController.PressKey(key)
		}
	}
}
