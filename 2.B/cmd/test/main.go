package main

import (
	"big2/pkg/card"
	"big2/pkg/game/component"
	"big2/pkg/game/round"
	"big2/pkg/player"
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {

	// player := player.NewPlayer(&player.PlayerOptions{
	// 	Core:   player.NewHumanPlayer(),
	// 	Writer: os.Stdout,
	// 	Reader: os.Stdin,
	// })

	// player.NamePlayer()
	// fmt.Println(player.Name)

	filePath := "tests/fullhouse.in"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	line = strings.TrimRight(line, "\n")
	cards := card.StringToCards(line)

	var inputs string
	for {
		other, err := reader.ReadString('\n')
		if err != nil {
			inputs += other
			break
		}
		inputs += other
	}

	buffer := bytes.NewBufferString(inputs)

	input := bufio.NewReader(buffer)

	game := &component.BigTwo{
		Players: []*player.Player{
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: os.Stdout,
				Reader: input,
			}),
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: os.Stdout,
				Reader: input,
			}),
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: os.Stdout,
				Reader: input,
			}),
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: os.Stdout,
				Reader: input,
			}),
		},
		Deck: component.NewDeck(&component.DeckOptions{
			Shuffle: false,
			Cards:   cards,
		}),
		Table: component.NewTable(),
	}

	game.
		AddRound(round.Deal).
		AddRound(round.NamePlayer).
		AddRound(round.DefaultPlay)

	game.Run()

	// for {
	// 	if err != nil {
	// 		break
	// 	}
	// }
}
