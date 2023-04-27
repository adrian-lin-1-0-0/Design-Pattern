package main

import (
	"bufio"
	"fmt"
	"os"

	"2.2.H/utils/bus"
)

func cmdServer(requestChan chan bus.Request) {
	for {
		req := <-requestChan
		fmt.Print(req.Cmd)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		res := bus.Response{Result: "Received: " + text}
		req.ReplyChan <- res
	}
}

func main() {
	requestChan := bus.NewBus()
	go cmdServer(requestChan)
	req := bus.Request{Cmd: "Input Name : ", ReplyChan: make(chan bus.Response)}
	requestChan <- req
	res := <-req.ReplyChan
	println(res.Result.(string))

}
