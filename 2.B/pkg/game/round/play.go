package round

import (
	"big2/pkg/card/patterns"
	"big2/pkg/game/component"
	"big2/pkg/notify/message"
	"fmt"
)

var DefaultPlay = NewPlay(patterns.CardPatternsFactory())

func NewPlay(cardPatternsChain *patterns.CardPatternsChain) func(*component.BigTwo) {
	return func(b *component.BigTwo) {

		playerCircular := NewPlayerCircular(b.Players)
		playerCount := len(b.Players)
		for playerCount > 0 {
			if p := playerCircular.GetPlayer(); hasClub3(p.HandCards()) {
				b.Table.TopPlayer = p
				break
			}
			playerCircular = playerCircular.Next()
			playerCount--
		}

		passCount := 0
		passLimit := len(b.Players) - 1
		isFirstRound := true

		for {
			p := playerCircular.GetPlayer()
			p.Begin()
			fmt.Fprintf(p.Writer, message.YourTurn, p.Name)
			for {
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
				p.Commit()
				break
			}

			for passLimit > passCount {
				playerCircular = playerCircular.Next()
				p := playerCircular.GetPlayer()
				fmt.Fprintf(p.Writer, message.YourTurn, p.Name)
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
					if len(p.HandCards()) == 0 {
						fmt.Fprintf(p.Writer, message.GameOver, p.Name)
						return
					}
					break
				}
			}
			b.Table.TopPlay = nil
			passCount = 0

		}

	}
}
