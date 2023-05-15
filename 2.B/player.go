package big2

type PlayerCore interface {
	NamePlayer(*Player)
	Play(*Player) []int
}

type Player struct {
	Name      string
	handCards HandCards
	core      PlayerCore
}

type PlayerOptions struct {
	Core PlayerCore
}

func (p *Player) Play() []Card {
	p.core.Play(p)
	return []Card{}
}

func (p *Player) Pass() {
	// TODO
}

func (p *Player) NamePlayer() {
	p.core.NamePlayer(p)
}

func (p *Player) DealtHandCards(c Card) {
	p.handCards.AddCard(c)
}

func (p *Player) BeginHand() {
	p.handCards.Begin()
}

func (p *Player) CommitHand() {
	p.handCards.Commit()
}

func (p *Player) RollbackHand() {
	p.handCards.Rollback()
}

func NewPlayer(opts *PlayerOptions) *Player {
	return &Player{core: opts.Core}
}
