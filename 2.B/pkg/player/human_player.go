package player

import (
	"fmt"
	"io"
	"os"

	"big2/pkg/notify/message"
)

type HumanPlayer struct {
	reader io.Reader
	writer io.Writer
}

type HumanPlayerOptions struct {
	Reader io.Reader
	Writer io.Writer
}

func (hp *HumanPlayer) NamePlayer(p *Player) {
	// fmt.Fprint(hp.writer, "Enter your name:")
	fmt.Fscanf(hp.reader, "%s", &p.Name)
}

func (hp *HumanPlayer) Play(p *Player) []int {
	fmt.Fprintf(hp.writer, message.YourTurn, p.Name)
	// TODO
	return []int{-1}
}

func NewHumanPlayer(opts *HumanPlayerOptions) *HumanPlayer {
	if opts == nil {
		opts = &HumanPlayerOptions{}
	}
	if opts.Reader == nil {
		opts.Reader = os.Stdin
	}
	if opts.Writer == nil {
		opts.Writer = os.Stdout
	}

	return &HumanPlayer{
		reader: opts.Reader,
		writer: opts.Writer,
	}
}
