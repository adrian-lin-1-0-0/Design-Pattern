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

		// Find the player with club 3
		playerCircular, err := playerCircular.Rotate(func(p *PlayerCircular) bool {
			return hasClub3(p.GetPlayer().HandCards())
		})

		if err != nil {
			panic(err)
		}

		passCount := 0
		passLimit := len(b.Players) - 1
		var cardPattern patterns.CardPattern
		isFirstRound := true

		for {
			p := playerCircular.GetPlayer()
			fmt.Fprintf(p.Writer, message.NewRound)
			fmt.Fprintf(p.Writer, message.YourTurn, p.Name)
			//topPlayer's turn
			for {
				//start player's transaction
				p.Begin()
				playerPlay := p.Play()
				if playerPlay == nil {
					goto TopPlayerPass
				}
				if isFirstRound {
					if !hasClub3(playerPlay) {
						goto TopPlayIllegalPlay
					}
					isFirstRound = false
				}
				cardPattern, err = cardPatternsChain.ToPattern(playerPlay)
				if err != nil {
					goto TopPlayIllegalPlay
				}

				goto TopPlayLegalPlay

			TopPlayerPass:
				fmt.Fprintf(p.Writer, message.CantPassInNewRound)
				p.Rollback()
				continue
			TopPlayIllegalPlay:
				fmt.Fprintf(p.Writer, message.IllegalPlay)
				p.Rollback()
				continue
			TopPlayLegalPlay:
				fmt.Fprintf(p.Writer, message.PlayerPlay, p.Name, cardPattern.GetName(), message.CardsToString(playerPlay))
				b.Table.TopPlay = cardPattern
				p.Commit()
				if len(p.HandCards()) == 0 {
					fmt.Fprintf(p.Writer, message.GameOver, p.Name)
					return
				}
				break
			}

			//other players' turn
			for passLimit > passCount {
				playerCircular = playerCircular.Next()
				p := playerCircular.GetPlayer()
				fmt.Fprintf(p.Writer, message.YourTurn, p.Name)

				for {
					p.Begin()
					playerPlay := p.Play()

					if playerPlay == nil {
						goto Pass
					}

					cardPattern, err = cardPatternsChain.ToPattern(playerPlay)
					if err != nil {
						goto IllegalPlay
					}

					if !cardPattern.GreaterThan(b.Table.TopPlay) {
						goto IllegalPlay
					}
					goto LegalPlay

				Pass:
					fmt.Fprintf(p.Writer, message.PlayerPass, p.Name)
					passCount++
					p.Commit()
					break
				IllegalPlay:
					fmt.Fprintf(p.Writer, message.IllegalPlay)
					p.Rollback()
					continue
				LegalPlay:
					fmt.Fprintf(p.Writer, message.PlayerPlay, p.Name, cardPattern.GetName(), message.CardsToString(playerPlay))
					b.Table.TopPlay = cardPattern
					b.Table.TopPlayer = p
					p.Commit()
					passCount = 0
					if len(p.HandCards()) == 0 {
						fmt.Fprintf(p.Writer, message.GameOver, p.Name)
						return
					}
					break
				}
			}

			//Clear the table
			passCount = 0
			b.Table.TopPlay = nil
			playerCircular = playerCircular.Next()
		}
	}
}
