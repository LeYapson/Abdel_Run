package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "ABDEL RUN!!!")
	rl.SetTargetFPS(60)

	rl.InitAudioDevice() // Initialise le module audio

	frameCounter := 0
	currentScreen := 0

	player := rl.Rectangle{
		X:      screenWidth / 8,
		Y:      screenHeight - 50.0,
		Width:  50,
		Height: 50,
	}

	isJumping := false
	velocity := float32(0.0)
	gravity := float32(1.0)
	jumpStrength := float32(-20.0)

	for !rl.WindowShouldClose() {
		switch currentScreen {
		case 0:
			rl.BeginDrawing()
			rl.ClearBackground(rl.Red)
			rl.EndDrawing()
			frameCounter++
			if frameCounter > 120 {
				currentScreen++
			}
		case 1:
			// Charger l'image de fond
			bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
			// Charger la musique de fond
			bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
			rl.PlayMusicStream(bgMusic)

			str := strconv.Itoa(currentScreen)
			rl.DrawText(str, 10, 0, 20, rl.Gray)

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

								rl.UnloadMusicStream(bgMusic) // Libérer la musique de la mémoire
								rl.CloseAudioDevice()         // Fermer le périphérique audio
								rl.CloseWindow()
								return
							case "Play":
								//rl.EndDrawing()
								//rl.UnloadMusicStream(bgMusic) // Libérer la musique de la mémoire
								//rl.CloseAudioDevice()
								currentScreen = 2
								rl.DrawText(str, 10, 0, 20, rl.Gray)
							}

							//if button.Text == "Quit" {

							//} else if button.Text == "Play" {
							//rl.EndDrawing()
							//rl.UnloadMusicStream(bgMusic) // Libérer la musique de la mémoire
							//rl.CloseAudioDevice()

							//rl.CloseWindow()

							//}
							// Gérer les autres boutons ici
						}
					}
					rl.DrawRectangleRec(button.Bounds, color)
					rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
				}
				rl.EndDrawing()
			}
		case 2:
			if rl.IsKeyPressed(rl.KeySpace) && !isJumping {
				isJumping = true
				velocity = jumpStrength
			}

			if isJumping {
				player.Y += velocity
				velocity += gravity
				if player.Y > screenHeight-50 {
					player.Y = screenHeight - 50
					isJumping = false
				}
			}
			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawText("PRESS SPACE to PAUSE MOVEMENT", 10, 0, 20, rl.Gray)
			rl.DrawRectangleRec(player, rl.Red)
			//rl.DrawTexture(texture, screenWidth/2-texture.width/2, screenHeight/2-texture.height/2, WHITE)
			rl.EndDrawing()
		}

	}
}
