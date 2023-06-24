package httpclient

type LoadBalancing struct {
	*UrlParser
	Roundrobin int
}

func NewLoadBalancing() *LoadBalancing {
	return &LoadBalancing{
		UrlParser: &UrlParser{},
	}
}

func (l *LoadBalancing) SendRequest(req HttpRequest) {
	url := req.Urls[(l.Roundrobin)%len(req.Urls)]
	l.Roundrobin += 1
	req.Urls = []string{url}
	l.Next.SendRequest(req)
}

func (l *LoadBalancing) SetNext(next HttpClient) *LoadBalancing {
	l.Next = next
	return l
}
