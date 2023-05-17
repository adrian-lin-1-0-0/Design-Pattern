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

		for {
			if p := playerCircular.GetPlayer(); hasClub3(p.HandCards()) {
				b.Table.TopPlayer = p
				break
			}
			playerCircular.Next()
			//TODO: handle no club3
		}

		passCount := 0
		passLimit := len(b.Players) - 1
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
