package httpclient

type HttpClient interface {
	SendRequest(HttpRequest)
}

type HttpRequest struct {
	Urls   []string
	Method string
}
