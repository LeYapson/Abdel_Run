package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

type Platform struct {
	Rect  rl.Rectangle
	Color rl.Color
}

type Obstacle struct {
	Rect  rl.Rectangle
	Color rl.Color
}

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

	//animFrames := int32(0)
	//p := &animFrames

	// Charger l'image de fond et la musique une seule fois
	bgLogo := rl.LoadTexture("../assets/images/logo.jpg")
	bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
	bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
	bgSettings := rl.LoadTexture("../assets/images/Settings.png")
	//gifAbdel := rl.LoadImageAnim("../assets/abdel_run.gif", p)

	// Créer des plateformes et des obstacles
	var platforms []Platform
	var obstacles []Obstacle

	lastPlatformHeight := float32(screenHeight) - 50.0 // la hauteur de départ de la première plateforme
	for i := 0; i < 5; i++ {
		platformWidth := float32(rl.GetRandomValue(200, 400)) // largeur aléatoire entre 200 et 400
		platformHeight := float32(20)                         // hauteur fixe pour une plateforme
		platformX := float32(rl.GetRandomValue(0, int32(screenWidth)-400))
		platformY := lastPlatformHeight - float32(rl.GetRandomValue(100, 150)) // espace entre 100 et 150

		platforms = append(platforms, Platform{Rect:  rl.NewRectangle(platformX, platformY, platformWidth, platformHeight),Color: rl.Blue,})

		// Ajouter un obstacle sur la plateforme
		obstacleWidth := float32(40)
		obstacleHeight := float32(40)
		obstacleX := platformX + float32(rl.GetRandomValue(10, int32(platformWidth)-50))
		obstacleY := platformY - obstacleHeight

		obstacles = append(obstacles, Obstacle{Rect:  rl.NewRectangle(obstacleX, obstacleY, obstacleWidth, obstacleHeight),Color: rl.Red})

		lastPlatformHeight = platformY
	}

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
							rl.UnloadMusicStream(bgMusic)
							rl.CloseAudioDevice()
							currentScreen = 2
						case "Settings":
							currentScreen = 3
						}

					}
				}
				rl.DrawRectangleRec(button.Bounds, color)
				rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
			}
			rl.EndDrawing()

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

			// Dessinez les plateformes et les obstacles
		for _, p := range platforms {
			rl.DrawRectangleRec(p.Rect, p.Color)
		}
		for _, o := range obstacles {
			rl.DrawRectangleRec(o.Rect, o.Color)
		}

			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawText("PRESS SPACE to JUMP", 10, 0, 20, rl.Gray)
			rl.DrawRectangleRec(player, rl.Red)
			rl.EndDrawing()
		case 3:
			rl.BeginDrawing()
			rl.DrawTexture(bgSettings, 0, 0, rl.White)
			rl.ClearBackground(rl.Green)
			rl.DrawText("Settings", screenWidth/2-100, 0, 50, rl.Black)

			buttons := []struct {
				Bounds rl.Rectangle
				Text   string
			}{
				{rl.NewRectangle(screenWidth/20, screenHeight/20, 150, 40), "Back"},
				{rl.NewRectangle(screenWidth-(150+screenWidth/20), screenHeight-(40+screenHeight/20), 150, 40), "Quit"},
			}

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
				rl.DrawRectangleRec(button.Bounds, color)
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
