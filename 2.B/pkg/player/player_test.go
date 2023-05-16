package player

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPlayer_NamePlayer(t *testing.T) {
	var input bytes.Buffer
	var output bytes.Buffer

	inputName := "Adrian"

	player := NewPlayer(&PlayerOptions{
		Core:   NewHumanPlayer(),
		Reader: &input,
		Writer: &output,
	})

	fmt.Fprintln(&input, inputName)
	player.NamePlayer()

	if player.Name != inputName {
		t.Errorf("player.Name = %s; want %s", player.Name, inputName)
	}

	player.Play()

}
