package ui

import (
	"strconv"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Settings(currentScreen int, screenWidth float32, screenHeight float32, fps int, bgMusic rl.Music, bgSettings rl.Texture2D) int {
	ratioArrondiRec := float32(0.5)
	segmentRec := int32(0)

	rl.BeginDrawing()
	rl.DrawTexture(bgSettings, 0, 0, rl.White)
	rl.ClearBackground(rl.Green)
	rl.DrawText("Settings", int32(screenWidth)/2-150, 0, 50, rl.Black)

	//Boutons back et quit dans settings
	buttons := []struct {
		Bounds rl.Rectangle
		Text   string
	}{
		{rl.NewRectangle(screenWidth/20, screenHeight/20, 150, 40), "Back"},
		{rl.NewRectangle(screenWidth-(150+screenWidth/20), screenHeight-(40+screenHeight/20), 150, 40), "Quit"},
		{rl.NewRectangle(screenWidth-220, screenHeight-360, 175, 40), "Change jump key"},
	}

	//Gestionnaire de FPS
	stringFPS := strconv.FormatInt(int64(fps), 10)
	rl.DrawText(stringFPS, int32(screenWidth)/2-150, int32(screenHeight)-420, 50, rl.Black)
	recSelectFPS := rl.NewRectangle(200, 230, 150, 30)
	fps = int(rg.SliderBar(recSelectFPS, "fps", "", float32(fps), 30, 144))
	rl.SetTargetFPS(int32(fps))

	//Affiche les FPS
	rl.DrawFPS(0, 0)

	//Création et fonctionnalité des boutons Back et Quit
	for _, button := range buttons {
		color := rl.Yellow
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.Bounds) {
			color = rl.DarkGray
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				switch button.Text {
				case "Back":
					currentScreen = 1
				case "Quit":
					currentScreen = 4
				case "Change jump key":

					// Afficher un message demandant d'appuyer sur une touche
					rl.DrawText("Appuyer sur une touche", 10, 10, 50, rl.Black)

					// Attendre qu'une touche soit pressée
					if rl.IsKeyPressed(rl.KeyZ) {
						//toucheSaut = rl.KeyZ
					}
				}
			}
		}
		rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
		rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
	}
	rl.EndDrawing()
	return currentScreen
}
