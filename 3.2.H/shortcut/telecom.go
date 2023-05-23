package shortcut

import "fmt"

type Telecom struct {
}

func (c *Telecom) Connect() {
	fmt.Println("The telecom has been turned on.")
}

func (c *Telecom) Disconnect() {
	fmt.Println("The telecom has been turned off.")
}

type ConnectTelecom struct {
	Telecom *Telecom
}

func (c ConnectTelecom) Do() {
	c.Telecom.Connect()
}

func (c ConnectTelecom) Undo() {
	c.Telecom.Disconnect()
}

type DisconnectTelecom struct {
	Telecom *Telecom
}

func (c DisconnectTelecom) Do() {
	c.Telecom.Disconnect()
}

func (c DisconnectTelecom) Undo() {
	c.Telecom.Connect()
}
