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
		Core: NewHumanPlayer(
			&HumanPlayerOptions{
				Reader: &input,
				Writer: &output,
			},
		),
	})

	fmt.Fprintln(&input, inputName)
	player.NamePlayer()

	if player.Name != inputName {
		t.Errorf("player.Name = %s; want %s", player.Name, inputName)
	}

	player.Play()

	expectedOutput := fmt.Sprintf("輪到%s了\n", player.Name)
	if output.String() != expectedOutput {
		t.Errorf("output = %s; want %s", output.String(), expectedOutput)
	}
}
