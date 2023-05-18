package round

import (
	"big2/pkg/card"
	"big2/pkg/player"
	"errors"
)

func hasClub3(cards []card.Card) bool {
	return hasCard(cards, card.Card{Suit: card.Clubs, Rank: card.Three})
}

func hasCard(cards []card.Card, card card.Card) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}

type PlayerCircular struct {
	next   *PlayerCircular
	Player *player.Player
	length int
}

func (p *PlayerCircular) Len() int {
	return p.length
}

func (p *PlayerCircular) Next() *PlayerCircular {
	p = p.next
	return p
}

func (p *PlayerCircular) GetPlayer() *player.Player {
	return p.Player
}

func (p *PlayerCircular) Rotate(f func(*PlayerCircular) bool) (*PlayerCircular, error) {
	count := 0
	limit := p.Len()
	for limit > count {
		if f(p) {
			return p, nil
		}
		if p.next == nil {
			return nil, errors.New("no player found")
		}
		p = p.next
	}
	return nil, errors.New("no player found")
}

func NewPlayerCircular(players []*player.Player) *PlayerCircular {
	var length = len(players)
	var head *PlayerCircular
	var tail *PlayerCircular
	for _, p := range players {
		if head == nil {
			head = &PlayerCircular{
				Player: p,
				length: length,
			}
			tail = head
			continue
		}
		tail.next = &PlayerCircular{
			Player: p,
		}
		tail = tail.next
	}
	tail.next = head
	return head
}
