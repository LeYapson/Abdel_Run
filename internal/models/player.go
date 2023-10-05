package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

type Player struct {
	hitbox rl.Rectangle
	health int
}

func NewPlayer(hitbox rl.Rectangle, health int) Player {
	p := Player{hitbox, health}
	return p
}
