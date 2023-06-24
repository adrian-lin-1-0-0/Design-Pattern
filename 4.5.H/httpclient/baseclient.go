package httpclient

import "fmt"

type BaseClient struct {
}

func (b *BaseClient) SendRequest(req HttpRequest) {
	fmt.Println(req.Urls[0])
}

func NewBaseClient() *BaseClient {
	return &BaseClient{}
}
