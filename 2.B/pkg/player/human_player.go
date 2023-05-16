package player

import (
	"fmt"
	"strconv"
	"strings"
)

type HumanPlayer struct {
}

func (hp *HumanPlayer) NamePlayer(p *Player) {
	// fmt.Fprint(hp.writer, "Enter your name:")
	fmt.Fscanf(p.Reader, "%s", &p.Name)
}

func (hp *HumanPlayer) Play(p *Player) []int {
	var idxLine string
	fmt.Fscanf(p.Reader, "%s", &idxLine)
	return strToIntArr(idxLine)
}

func NewHumanPlayer() *HumanPlayer {
	return &HumanPlayer{}
}

func strToIntArr(str string) []int {
	idxStrList := strings.Split(str, "")
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
