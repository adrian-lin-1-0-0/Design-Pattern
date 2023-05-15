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
	play := p.core.Play(p)
	if len(play) == 1 && play[0] == -1 {
		return nil
	}

	playCards := []Card{}
	handCards := p.handCards.GetCards()
	for _, handIdx := range play {
		playCards = append(playCards, handCards[handIdx])
	}

	//TODO
	//remove play cards from hand cards

	return playCards
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

func (p *Player) Begin() {
	p.handCards.Begin()
}

func (p *Player) Commit() {
	p.handCards.Commit()
}

func (p *Player) Rollback() {
	p.handCards.Rollback()
}

func NewPlayer(opts *PlayerOptions) *Player {
	return &Player{core: opts.Core}
}
