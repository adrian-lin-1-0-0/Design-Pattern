package player

import "big2/pkg/card"

type HandCards struct {
	data   []card.Card
	backup []card.Card
	set    func([]card.Card)
	get    func() []card.Card
}

func NewHandCards() *HandCards {
	h := &HandCards{
		data:   []card.Card{},
		backup: []card.Card{},
	}
	h.useData()
	return h
}

func (h *HandCards) SetCards(cards []card.Card) {
	h.set(cards)
}

func (h *HandCards) GetCards() []card.Card {
	return h.get()
}

func (h *HandCards) AddCard(card card.Card) {
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

func (h *HandCards) setData(data []card.Card) {
	h.data = append([]card.Card{}, data...)
}

func (h *HandCards) setBackup(data []card.Card) {
	h.backup = append([]card.Card{}, data...)
}

func (h *HandCards) useBackup() {
	h.set = h.setBackup
	h.get = h.getBackup
}

func (h *HandCards) useData() {
	h.set = h.setData
	h.get = h.getData
}

func (h *HandCards) getData() []card.Card {
	return h.data
}

func (h *HandCards) getBackup() []card.Card {
	return h.backup
}

func (h *HandCards) clearBackup() {
	h.setBackup([]card.Card{})
}
