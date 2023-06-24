package httpclient

type BlockList struct {
	*BaseClient
	list []string
}

func NewBlockList() *BlockList {
	return &BlockList{
		list: []string{""},
	}
}

func (b *BlockList) CheckBlock(url string) bool {
	for _, blockUrl := range b.list {
		if url == blockUrl {
			return true
		}
	}
	return false
}

func (b *BlockList) SendRequest(req HttpRequest) {
	if b.CheckBlock(req.Urls[0]) {
		panic("blocked")
	}
	req.Urls = []string{req.Urls[0]}
	b.BaseClient.SendRequest(req)
}
