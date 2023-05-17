package round

import (
	"big2/pkg/card"
	"big2/pkg/card/patterns"
	"big2/pkg/game"
	"big2/pkg/notify/message"
	"big2/pkg/player"
	"fmt"
)

func NewPlay(cardPatternsChain *patterns.CardPatternsChain) func(*game.BigTwo) {
	return func(b *game.BigTwo) {

		passCount := 0
		passLimit := len(b.Players) - 1
		playerCircular := NewPlayerCircular(b.Players)

		for {
			if p := playerCircular.GetPlayer(); hasClub3(p.HandCards()) {
				b.Table.TopPlayer = p
				break
			}
			playerCircular.Next()
			//TODO: handle no club3
		}

		isFirstRound := true

		for {
			p := playerCircular.GetPlayer()
			p.Begin()

			{
				topPlay := p.Play()

				if topPlay == nil {
					fmt.Fprintf(p.Writer, message.CantPassInNewRound)
					p.Rollback()
					continue
				}
				if isFirstRound {
					if !hasClub3(topPlay) {
						fmt.Fprintf(p.Writer, message.IllegalPlay)
						p.Rollback()
						continue
					}
					isFirstRound = false
				}
				cardPattern, err := cardPatternsChain.ToPattern(topPlay)
				if err != nil {
					fmt.Fprintf(p.Writer, message.IllegalPlay)
					p.Rollback()
					continue
				}

				b.Table.TopPlay = cardPattern
			}

			p.Commit()

			for passLimit > passCount {
				playerCircular.Next()
				p := playerCircular.GetPlayer()
				p.Begin()

				for {
					topPlay := p.Play()

					if topPlay == nil {
						fmt.Fprintf(p.Writer, message.PlayerPass, p.Name)
						passCount++
						p.Commit()
						break
					}

					cardPattern, err := cardPatternsChain.ToPattern(topPlay)
					if err != nil {
						fmt.Fprintf(p.Writer, message.IllegalPlay)
						p.Rollback()
						continue
					}

					if !cardPattern.GreaterThan(b.Table.TopPlay) {
						fmt.Fprintf(p.Writer, message.IllegalPlay)
						p.Rollback()
						continue
					}

					b.Table.TopPlay = cardPattern
					b.Table.TopPlayer = p
					passCount = 0
					if len(p.HandCards()) == 0 {
						fmt.Fprintf(p.Writer, message.GameOver, p.Name)
						return
					}
					break
				}
			}
			b.Table.TopPlay = nil

		}

	}
}

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
