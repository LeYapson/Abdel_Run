package player

type Position struct {
	X, Y int
}

var (
	Left   = 0
	Center = 1
	Right  = 2
)

type Player struct {
	Pos    Position
	Moving bool
}

func New() *Player {
	return &Player{}
}

func (p *Player) MoveLeft() {
	if p.Moving {
		return
	}
	if p.Pos.X > Left {
		p.Pos.X--
		p.Moving = true
	}
}

func (p *Player) MoveRight() {
	if p.Moving {
		return
	}
	if p.Pos.X < Right {
		p.Pos.X++
		p.Moving = true
	}
}
