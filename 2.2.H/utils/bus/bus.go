package bus

type Request struct {
	Cmd       string
	Data      interface{}
	ReplyChan chan Response
}

type Response struct {
	Result interface{}
}

func NewBus() chan Request {
	requestChan := make(chan Request)
	return requestChan
}
