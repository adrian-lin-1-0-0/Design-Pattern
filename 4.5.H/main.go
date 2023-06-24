package main

import (
	"4.5.H/httpclient"
)

func main() {
	FakeHttpClient := httpclient.NewLoadBalancing().
		SetNext(httpclient.NewServiceDiscovery().
			SetNext(httpclient.NewBlockList()))
	FakeHttpClient.SendRequest(httpclient.HttpRequest{
		Urls:   []string{"http://waterballsa.tw/mail"},
		Method: "GET",
	})
	// http://35.0.0.1/mail

	FakeHttpClient2 := httpclient.NewServiceDiscovery().
		SetNext(httpclient.NewLoadBalancing().
			SetNext(httpclient.NewBlockList()))

	FakeHttpClient2.SendRequest(httpclient.HttpRequest{
		Urls:   []string{"http://waterballsa.tw/mail"},
		Method: "GET",
	})
	// http://35.0.0.1/mail
}
