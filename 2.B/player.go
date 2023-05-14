package big2

type PlayerCore interface {
	NamePlayer(*Player)
	Play(*Player) []int
}

type Player struct {
	Name      string
	handCards []Card
	core      PlayerCore
}

type PlayerOptions struct {
	Core PlayerCore
}

func (p *Player) Play() {
	p.core.Play(p)[]int
}

func (p *Player) Pass() {
	// TODO
}

func (p *Player) NamePlayer() {
	p.core.NamePlayer(p)
}

func (p *Player) DealtHandCards(c Card) {
	p.handCards = append(p.handCards, c)
}

func NewPlayer(opts *PlayerOptions) *Player {
	return &Player{core: opts.Core}
}
