package shortcut

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Terminal struct {
	MainController           *MainController
	ConnectTelecom           ConnectTelecom
	DisconnectTelecom        DisconnectTelecom
	MoveTankBackward         MoveTankBackward
	MoveTankForward          MoveTankForward
	ResetMainControlKeyboard ResetMainControlKeyboard
	options                  []Command
}

func NewTerminal() *Terminal {
	telecom := &Telecom{}
	tank := &Tank{}
	terminal := &Terminal{
		MainController: NewMainController(),
		ResetMainControlKeyboard: ResetMainControlKeyboard{
			MainController: &MainController{},
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

	terminal.options = []Command{
		terminal.MoveTankBackward,
		terminal.MoveTankBackward,
		terminal.ConnectTelecom,
		terminal.DisconnectTelecom,
		terminal.ResetMainControlKeyboard,
	}
	return terminal
}

func (c *Terminal) SetShortcut() {
	var wantMacro, key rune
	fmt.Print("設置巨集指令 (y/n): ")
	fmt.Scanf("%c\n", &wantMacro)
	fmt.Print("Key: ")
	fmt.Scanf("%c\n", &key)

	if wantMacro == rune('y') {
		fmt.Printf("要將哪些指令設置成快捷鍵 %c 的巨集（輸入多個數字，以空白隔開）:\n", key)
		for i, v := range c.options {
			fmt.Printf("(%d) %s\n", i, v.Name())
		}
		var commands string
		in := bufio.NewReader(os.Stdin)
		commands, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		setMacro(commands, key, c)
	} else {
		fmt.Printf("要將哪一道指令設置到快捷鍵 %c 上: \n", key)
		for i, v := range c.options {
			fmt.Printf("(%d) %s\n", i, v.Name())
		}
		var command rune
		fmt.Scanf("%c\n", &command)
		if idx := runeToOptionsIdx(command, c.options); idx != -1 {
			c.MainController.KeyBind(key, c.options[idx])
		}
	}
}

func (c *Terminal) Run() {
	var key rune
	for {
		for k, v := range c.MainController.keyBindings {
			fmt.Printf("%c: %s\n", k, v.Name())
		}
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

func setMacro(input string, key rune, c *Terminal) {
	commands := strings.Split(string(input), " ")
	macro := NewMacro()
	for _, v := range commands {
		if idx := runeToOptionsIdx(rune(v[0]), c.options); idx != -1 {
			macro.Add(c.options[idx])
		}
	}
	c.MainController.KeyBind(key, macro)
}

func runeToOptionsIdx(r rune, options []Command) int {
	idx := int(r - '0')
	if idx >= len(options) {
		fmt.Println("指令錯誤")
		return -1
	}
	return idx
}
