package player

import (
	"strconv"
	"strings"
)

type HumanPlayer struct {
}

func (hp *HumanPlayer) NamePlayer(p *Player) {
	// fmt.Fprint(hp.writer, "Enter your name:")
	p.Name = readLine(p.Reader)
}

func (hp *HumanPlayer) Play(p *Player) []int {
	return strToIntArr(readLine(p.Reader))
}

func NewHumanPlayer() *HumanPlayer {
	return &HumanPlayer{}
}

func strToIntArr(str string) []int {
	idxStrList := strings.Split(str, " ")
	intArr := []int{}
	for _, idxStr := range idxStrList {
		idx, err := strconv.ParseInt(idxStr, 10, 8)
		if err != nil {
			continue
		}
		intArr = append(intArr, int(idx))
	}
	return intArr
}
