package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"

	"github.com/LeYapson/Abdel_Run/internal/player"
	ui "github.com/LeYapson/Abdel_Run/internal/ui"
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	gifImages    []*ebiten.Image
	currentFrame int
	counter      int
	delays       []int
}

var abdel *player.Player

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		abdel.MoveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		abdel.MoveRight()
	} else {
		abdel.Moving = false
	}

	// Update GIF animation
	g.counter++
	if g.counter > g.delays[g.currentFrame] {
		g.currentFrame = (g.currentFrame + 1) % len(g.gifImages)
		g.counter = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

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
	screen.DrawImage(g.gifImages[g.currentFrame], opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1024, 680
}

func loadGIF(path string) ([]*ebiten.Image, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		return nil, nil, err
	}

	images := make([]*ebiten.Image, len(g.Image))
	delays := make([]int, len(g.Image))
	previous := image.NewRGBA(g.Image[0].Bounds())

	for i, img := range g.Image {
		draw.Draw(previous, img.Bounds(), img, image.Point{}, draw.Over)
		rgba := image.NewRGBA(g.Image[0].Bounds())
		draw.Draw(rgba, g.Image[0].Bounds(), previous, image.Point{}, draw.Src)
		images[i] = ebiten.NewImageFromImage(rgba)
		delays[i] = g.Delay[i]
	}
	return images, delays, nil
}

func main() {
	ui.TitleScreen()
}
