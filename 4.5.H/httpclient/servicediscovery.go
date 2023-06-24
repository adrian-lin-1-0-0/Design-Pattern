package httpclient

import (
	"strings"
)

type ServiceDiscovery struct {
	*UrlParser
	dns map[string][]string
}

func NewServiceDiscovery() *ServiceDiscovery {
	dns := make(map[string][]string)
	dns["waterballsa.tw"] = []string{"35.0.0.1", "35.0.0.2", "35.0.0.3"}
	return &ServiceDiscovery{
		UrlParser: &UrlParser{},
		dns:       dns,
	}
}

func (s *ServiceDiscovery) SendRequest(req HttpRequest) {
	url := req.Urls[0]
	domain := strings.Split(url, "/")[2]
	ips := s.dns[domain]

	newIps := []string{}
	for _, ip := range ips {
		newIp := strings.Replace(url, domain, ip, 1)
		newIps = append(newIps, newIp)
	}
	req.Urls = newIps
	s.Next.SendRequest(req)
}

func (s *ServiceDiscovery) SetNext(next HttpClient) *ServiceDiscovery {
	s.Next = next
	return s
}
