package main

import (
	"strconv"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

var fps = int32(60)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "ABDEL RUN!!!")
	rl.SetTargetFPS(fps)
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
	ratioArrondiRec := float32(0.5)
	segmentRec := int32(0)

	//animFrames := int32(0)
	//p := &animFrames

	// Charger l'image de fond et la musique une seule fois
	bgLogo := rl.LoadTexture("../assets/images/logo.jpg")
	bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
	bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
	bgSettings := rl.LoadTexture("../assets/images/Settings.png")
	//gifAbdel := rl.LoadImageAnim("../assets/abdel_run.gif", p)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(bgMusic)

		switch currentScreen {
		case 0:
			rl.BeginDrawing()
			rl.DrawTexture(bgLogo, 0, 0, rl.White)
			rl.EndDrawing()
			frameCounter++
			if frameCounter > 120 {
				currentScreen++
				rl.PlayMusicStream(bgMusic)
			}

		case 1:
			rl.BeginDrawing()
			rl.DrawTexture(bgImage, 0, 0, rl.White)

			if !rl.IsMusicStreamPlaying(bgMusic) {
				//rl.UpdateMusicStream(bgMusic)
				//rl.ResumeMusicStream(bgMusic)
				rl.UpdateMusicStream(bgMusic)
				rl.PlayMusicStream(bgMusic)
			}
			buttons := []struct {
				Bounds rl.Rectangle
				Text   string
			}{
				{rl.NewRectangle(screenWidth-325, screenHeight/2-40, 150, 40), "Play"},
				{rl.NewRectangle(screenWidth-325, screenHeight/2+10, 150, 40), "Settings"},
				{rl.NewRectangle(screenWidth-325, screenHeight/2+60, 150, 40), "Quit"},
			}

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
							rl.StopMusicStream(bgMusic)
							currentScreen = 2
						case "Settings":
							currentScreen = 3
						}

					}
				}
				rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
				rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
			}
			rl.EndDrawing()

		case 2:

			//Saut du personnage
			if rl.IsKeyPressed(rl.KeySpace) && !isJumping {
				isJumping = true
				velocity = jumpStrength
			}

			//Retombée du personnage
			if isJumping {
				player.Y += velocity
				velocity += gravity
				if player.Y > screenHeight-player.Height {
					player.Y = screenHeight - player.Height
					isJumping = false
				}
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawText("PRESS SPACE to JUMP", 10, 0, 20, rl.Gray)
			rl.DrawRectangleRec(player, rl.Red)

			//Création du bouton back
			buttons := []struct {
				Bounds rl.Rectangle
				Text   string
			}{
				{rl.NewRectangle(screenWidth/20, screenHeight/20, 150, 40), "Back"},
			}

			//Fonctionnalité du bouton
			for _, button := range buttons {
				color := rl.Yellow
				if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.Bounds) {
					color = rl.DarkGray
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						switch button.Text {
						case "Back":
							currentScreen = 1
						}
					}
				}

				//Affichage du bouton
				rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
				rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
			}

			rl.EndDrawing()

		case 3:
			rl.BeginDrawing()
			rl.DrawTexture(bgSettings, 0, 0, rl.White)
			rl.ClearBackground(rl.Green)
			rl.DrawText("Settings", screenWidth/2-150, 0, 50, rl.Black)

			//Boutons back et quit dans settings
			buttons := []struct {
				Bounds rl.Rectangle
				Text   string
			}{
				{rl.NewRectangle(screenWidth/20, screenHeight/20, 150, 40), "Back"},
				{rl.NewRectangle(screenWidth-(150+screenWidth/20), screenHeight-(40+screenHeight/20), 150, 40), "Quit"},
			}

			//Gestionnaire de FPS
			stringFPS := strconv.FormatInt(int64(fps), 10)
			rl.DrawText(stringFPS, screenWidth/2-100, screenHeight-50, 50, rl.Black)
			recSelectFPS := rl.NewRectangle(300, 250, 105, 20)
			fps = int32(rg.SliderBar(recSelectFPS, "fps", "", float32(fps), 30, 144))
			rl.SetTargetFPS(fps)

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
							rl.UnloadMusicStream(bgMusic)
							rl.CloseAudioDevice()
							rl.CloseWindow()
							return
						}

					}
				}
				rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
				rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
			}
			rl.EndDrawing()
		}

	}

	rl.UnloadTexture(bgImage)
	rl.UnloadMusicStream(bgMusic)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
