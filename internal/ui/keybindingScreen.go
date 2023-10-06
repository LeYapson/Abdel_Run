package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var toucheSaut int32

func KeybindingScreen(currentScreen int, ratioArrondiRec float32, segmentRec int32, bgLogo rl.Texture2D, bgMusic rl.Music) (int, int32) {
	// ratioArrondiRec := float32(0.5)
	// segmentRec := int32(0)
	rl.ClearBackground(rl.Green)
	rl.DrawText("Press a key to change jump key", 100, 100, 50, rl.Black)

	key := rl.GetKeyPressed()
	fmt.Println(key)
	if key != 0 {
		toucheSaut = key
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.Green)

	//Bouton back
	buttons := []struct {
		Bounds rl.Rectangle
		Text   string
	}{
		{rl.NewRectangle(screenWidth/20, screenHeight/20, 150, 40), "Back"},
	}

	//Création et fonctionnalité des boutons Back et Quit
	for _, button := range buttons {
		color := rl.Yellow
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.Bounds) {
			color = rl.DarkGray
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				switch button.Text {
				case "Back":
					currentScreen = 3
				}
			}
		}
		rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
		rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
	}
	rl.EndDrawing()
	return currentScreen, toucheSaut
}
