package httpclient

type UrlParser struct {
	Next HttpClient
}

func NewUrlParser() *UrlParser {
	return &UrlParser{}
}

func (u *UrlParser) SendRequest(req HttpRequest) {
	panic("implement me")
}

func (u *UrlParser) SetNext(next HttpClient) *UrlParser {
	u.Next = next
	return u
}
