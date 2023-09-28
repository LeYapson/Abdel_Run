// Le paquet principal pour l'exécutable du jeu "Abdel_Run".
package main

// Importation des bibliothèques nécessaires.
import (
	"image/color"
	"log"

	"github.com/LeYapson/Abdel_Run/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Définition du type Game, qui est une struct vide.
// Elle est utilisée pour implémenter l'interface ebiten.Game.
type Game struct{}

// Déclaration d'une variable globale "abdel" pour représenter le joueur.
var abdel *player.Player

// La méthode Update est appelée à chaque frame pour gérer les mises à jour de jeu.
func (g *Game) Update() error {
	// Vérifie si la touche de gauche est enfoncée.
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		abdel.MoveLeft() // Déplace Abdel vers la gauche.
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		abdel.MoveRight() // Déplace Abdel vers la droite.
	} else {
		// Si aucune des touches n'est enfoncée, réinitialise l'état de déplacement d'Abdel.
		abdel.Moving = false
	}
	return nil
}

// La méthode Draw est appelée à chaque frame pour gérer le rendu graphique.
func (g *Game) Draw(screen *ebiten.Image) {
	// Définit la couleur de fond de l'écran à noir.
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	ebitenutil.DebugPrintAt(screen, "Hello World", 0, 150)

	// Crée une image pour représenter Abdel. Ici, il est représenté par un rectangle blanc.
	abdelRect := ebiten.NewImage(50, 50)
	abdelRect.Fill(color.White)

	// Calcule la position X d'Abdel en fonction de sa position actuelle.
	var abdelX int
	switch abdel.Pos {
	case player.Left:
		abdelX = 50
	case player.Center:
		abdelX = 225
	case player.Right:
		abdelX = 400
	}

	// Prépare les options de dessin pour Abdel.
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(abdelX), 500)

	// Dessine l'image d'Abdel sur l'écran.
	screen.DrawImage(abdelRect, opts)
}

// La méthode Layout définit la taille de la fenêtre.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 500, 600
}

// La fonction main est le point d'entrée du programme.
func main() {
	// Initialise Abdel.
	abdel = player.New()

	// Définit la taille de la fenêtre et son titre.
	ebiten.SetWindowSize(1024, 600)
	ebiten.SetWindowTitle("Abdel_Run")

	// Crée une nouvelle instance du jeu et la lance.
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		// Si une erreur se produit lors de l'exécution du jeu, elle est enregistrée et le programme est arrêté.
		log.Fatal(err)
	}
}
