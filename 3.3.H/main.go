package main

import (
	"os"
	"os/exec"

	"3.3.H/domain"
)

func main() {
	m := domain.DefaultMap()
	for {
		// clear()
		m.ShowMap()
		m.TakeTurn()
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
