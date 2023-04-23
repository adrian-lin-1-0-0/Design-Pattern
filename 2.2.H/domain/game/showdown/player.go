package showdown

type PlayerCore struct {
	point int
}

func (playerCore *PlayerCore) AddPoint() {
	playerCore.point++
}

func (playerCore *PlayerCore) GetPoint() int {
	return playerCore.point
}
