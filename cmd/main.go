package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math/rand"
	"os"

	"github.com/LeYapson/Abdel_Run/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gifImages            []*ebiten.Image
	currentFrame         int
	counter              int
	delays               []int
	obstacles            []*Obstacle
	obstacleSpawnCounter int
}

type Obstacle struct {
	X, Y  float64
	Type  int
	Lane  player.Position
	Image *ebiten.Image
}

var abdel *player.Player

const obstacleSpawnInterval = 60 // par exemple, un nouvel obstacle toutes les 60 frames

func (g *Game) spawnObstacle() {
	// Pour l'instant, nous ajoutons un obstacle simple, à améliorer plus tard
	path := rand.Intn(3) // 0: gauche, 1: centre, 2: droite

	var x float64
	switch path {
	case 0:
		x = 50
	case 1:
		x = 295
	case 2:
		x = 540
	}

	obstacle := &Obstacle{
		X:     x,
		Y:     0, // en haut de l'écran
		Type:  1,
		Image: ebiten.NewImage(50, 50), // une image simple pour l'instant
	}
	obstacle.Image.Fill(color.RGBA{255, 0, 0, 255}) // rempli en rouge

	g.obstacles = append(g.obstacles, obstacle)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		abdel.MoveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		abdel.MoveRight()
	} else {
		abdel.Moving = false
	}

	g.obstacleSpawnCounter++
	if g.obstacleSpawnCounter >= obstacleSpawnInterval {
		g.spawnObstacle()
		g.obstacleSpawnCounter = 0

		// Update GIF animation
		g.counter++
		if g.counter > g.delays[g.currentFrame] {
			g.currentFrame = (g.currentFrame + 1) % len(g.gifImages)
			g.counter = 0
		}
	}

	for i, obs := range g.obstacles {
        obs.Y += 10 // définissez une vitesse constante pour les obstacles
        if obs.Y > 600 {
            g.obstacles = append(g.obstacles[:i], g.obstacles[i+1:]...)
        }
	}

	

	for _, obs := range g.obstacles {
        if isColliding(obs, abdel) {
            // Collision détectée! Traitez ici (arrêtez le jeu, réduisez les vies, etc.)
        }
    }

	return nil
}





func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	var abdelX int
	switch abdel.Pos {
	case player.Position.Left:
		abdelX = 25
	case player.Center:
		abdelX = 225
	case player.Right:
		abdelX = 425
	}

	log.Println("Abdel position", abdelX)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(abdelX), 500)
	screen.DrawImage(g.gifImages[g.currentFrame], opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 500, 600
}

func isColliding(obs *Obstacle, ply *player.Player) bool {
    // Récupérer les dimensions et positions des rectangles représentant le joueur et l'obstacle.
    playerRect := image.Rect(int(ply.Pos.X), int(ply.Pos.Y), int(ply.Pos.X)+64, int(ply.Pos.Y)+64)
    obstacleRect := image.Rect(int(obs.X), int(obs.Y), int(obs.X)+obs.Width, int(obs.Y)+obs.Height)

    // Vérifier si les rectangles se chevauchent.
    return playerRect.Overlaps(obstacleRect)
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
	abdel = player.New()

	// Load the GIF for Abdel
	gifImages, delays, err := loadGIF("../assets/images/abdel_run.gif")
	if err != nil {
		log.Fatalf("Error loading Abdel's GIF: %v", err)
	}

	ebiten.SetWindowSize(1024, 600)
	ebiten.SetWindowTitle("Abdel_Run")
	game := &Game{
		gifImages: gifImages,
		delays:    delays,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
