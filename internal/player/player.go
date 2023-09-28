package player

type Position int

const (
	Left   Position = iota
	Center
	Right
)

type Player struct {
	Pos Position
	Moving bool
}


func New() *Player {
	return &Player{
		Pos: Center,
	}
}

func (p *Player) MoveLeft() {
	if p.Moving {
		return
	}
	if p.Pos > Left {
		p.Pos--
		p.Moving = true
	}
}

func (p *Player) MoveRight() {
	if p.Moving {
		return
	}
	if p.Pos < Right {
		p.Pos++
		p.Moving = true
	}
}

