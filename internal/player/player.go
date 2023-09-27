package player

type Position int

const (
	Left   Position = iota
	Center
	Right
)

type Player struct {
	Pos Position
}

func New() *Player {
	return &Player{
		Pos: Center,
	}
}

func (p *Player) MoveLeft() {
	if p.Pos > Left {
		p.Pos--
	}
}

func (p *Player) MoveRight() {
	if p.Pos < Right {
		p.Pos++
	}
}
