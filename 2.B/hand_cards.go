package big2

type HandCards struct {
	data   []Card
	backup []Card
	set    func([]Card)
	get    func() []Card
}

func NewHandCards() *HandCards {
	h := &HandCards{
		data:   []Card{},
		backup: []Card{},
	}
	h.useData()
	return h
}

func (h *HandCards) SetCards(cards []Card) {
	h.set(cards)
}

func (h *HandCards) GetCards() []Card {
	return h.get()
}

func (h *HandCards) AddCard(card Card) {
	h.set(append(h.get(), card))
}

func (h *HandCards) Begin() {
	h.setBackup(h.data)
	h.useBackup()
}

func (h *HandCards) Commit() {
	h.setData(h.backup)
	h.clearBackup()
	h.useData()
}

func (h *HandCards) Rollback() {
	h.clearBackup()
	h.useData()
}

func (h *HandCards) setData(data []Card) {
	h.data = append([]Card{}, data...)
}

func (h *HandCards) setBackup(data []Card) {
	h.backup = append([]Card{}, data...)
}

func (h *HandCards) useBackup() {
	h.set = h.setBackup
	h.get = h.getBackup
}

func (h *HandCards) useData() {
	h.set = h.setData
	h.get = h.getData
}

func (h *HandCards) getData() []Card {
	return h.data
}

func (h *HandCards) getBackup() []Card {
	return h.backup
}

func (h *HandCards) clearBackup() {
	h.setBackup([]Card{})
}
