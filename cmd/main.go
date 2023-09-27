package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/LeYapson/Abdel_Run/internal/player"
)

type Game struct{}

var abdel *player.Player

func (g *Game) Update() error {
	// Gérer les entrées pour déplacer Abdel
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		abdel.MoveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		abdel.MoveRight()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Remplir l'écran avec une couleur de fond (ici noir)
	screen.Fill(color.Black)

	// Dessiner Abdel en fonction de sa position
	abdelRect := ebiten.NewImage(50, 50)
	abdelRect.Fill(color.White)  // Pour simplifier, Abdel est un rectangle blanc

	var abdelX int
	switch abdel.Pos {
	case player.Left:
		abdelX = 50
	case player.Center:
		abdelX = 295
	case player.Right:
		abdelX = 540
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(abdelX), 215)
	screen.DrawImage(abdelRect, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // Taille initiale de la fenêtre
}

func main() {
	abdel = player.New()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Abdel_Run")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
