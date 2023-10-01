package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)
 
func TitleScreen() {
	var currentScreen = 1
	rl.InitWindow(screenWidth, screenHeight, "ABDEL RUN!!!")
	rl.SetTargetFPS(60)

	rl.InitAudioDevice() // Initialise le module audio

	// Charger l'image de fond
	bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
	// Charger la musique de fond
	bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
	rl.PlayMusicStream(bgMusic)

	buttons := []struct {
		Bounds rl.Rectangle
		Text   string
	}{
		{rl.NewRectangle(screenWidth-325, screenHeight/2-40, 150, 40), "Play"},
		{rl.NewRectangle(screenWidth-325, screenHeight/2+10, 150, 40), "Settings"},
		{rl.NewRectangle(screenWidth-325, screenHeight/2+60, 150, 40), "Quit"},
	}

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(bgMusic) // Mettre à jour le flux de la musique

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		// Dessiner l'image de fond
		rl.DrawTexture(bgImage, 0, 0, rl.White)

		for _, button := range buttons {
			color := rl.Yellow
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.Bounds) {
				color = rl.DarkGray
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					switch button.Text {
					case "Quit":
						rl.UnloadMusicStream(bgMusic)
						rl.CloseAudioDevice()
						rl.CloseWindow()
						return
					case "Play":
						rl.UnloadMusicStream(bgMusic)
						rl.CloseAudioDevice()
						currentScreen += 1
					case "Settings":
						currentScreen += 2
					}

				}
			}
			rl.DrawRectangleRec(button.Bounds, color)
			rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
		}
		rl.EndDrawing()
}
}