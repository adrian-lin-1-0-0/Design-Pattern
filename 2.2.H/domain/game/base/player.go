package base

type (
	IPlayerCore interface {
		NameSelf(IPlayer)
		TakeTurn(IPlayer) ICard
	}

	IPlayer interface {
		AddCard(ICard)
		NameSelf()
		TakeTurn() ICard
		ShowHand() IHand
		SetName(string)
		Name() string
	}

	BasePlayer struct {
		PlayerCore IPlayerCore
		name       string
		Hand       BaseHand
	}
)

func (player *BasePlayer) Name() string {
	return player.name
}

func (player *BasePlayer) SetName(name string) {
	player.name = name
}

func (player *BasePlayer) ShowHand() IHand {
	return &player.Hand
}

func (player *BasePlayer) AddCard(card ICard) {
	player.Hand.AddCard(card)
}

func (player *BasePlayer) TakeTurn() ICard {
	return player.PlayerCore.TakeTurn(player)
}

func (player *BasePlayer) NameSelf() {
	player.PlayerCore.NameSelf(player)
}

func NewBasePlayer(core IPlayerCore) *BasePlayer {
	return &BasePlayer{
		PlayerCore: core,
	}
}
