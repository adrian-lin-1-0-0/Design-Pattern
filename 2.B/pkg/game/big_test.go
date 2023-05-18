package game

import (
	"big2/pkg/card"
	"big2/pkg/game/component"
	"big2/pkg/game/round"
	"big2/pkg/player"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func fileToString(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func runGameWithFile(fileName string) string {
	filePath := fmt.Sprintf("../../tests/%s.in", fileName)
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
	writer := &strings.Builder{}

	game := &component.BigTwo{
		Players: []*player.Player{
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: writer,
				Reader: input,
			}),
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: writer,
				Reader: input,
			}),
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: writer,
				Reader: input,
			}),
			player.NewPlayer(&player.PlayerOptions{
				Core:   player.NewHumanPlayer(),
				Writer: writer,
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
	return writer.String()
}

func TestGame(t *testing.T) {

	tests := []struct {
		fileName string
	}{
		{
			fileName: "always-play-first-card",
		},
		{
			fileName: "fullhouse",
		},
		{
			fileName: "normal-no-error-play2",
		},
		{
			fileName: "normal-no-error-play1",
		},
		{
			fileName: "straight",
		},
		{
			fileName: "illegal-actions",
		},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			want := fileToString(fmt.Sprintf("../../tests/%s.out", tt.fileName))
			if got := runGameWithFile(tt.fileName); got != want {
				t.Errorf("Game() = %v, want %v", got, want)
			}
		})
	}

}
