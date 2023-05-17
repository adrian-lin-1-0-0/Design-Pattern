package round

import (
	"big2/pkg/card"
	"big2/pkg/player"
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

type playerCircular struct {
	next   *playerCircular
	Player *player.Player
}

func (p *playerCircular) Next() *playerCircular {
	p = p.next
	return p
}

func (p *playerCircular) GetPlayer() *player.Player {
	return p.Player
}

func NewPlayerCircular(players []*player.Player) *playerCircular {
	var head *playerCircular
	var tail *playerCircular
	for _, p := range players {
		if head == nil {
			head = &playerCircular{
				Player: p,
			}
			tail = head
			continue
		}
		tail.next = &playerCircular{
			Player: p,
		}
		tail = tail.next
	}
	tail.next = head
	return head
}
