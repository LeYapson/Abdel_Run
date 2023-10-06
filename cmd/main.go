package main

import (
	"fmt"
	"strconv"

	ui "github.com/LeYapson/Abdel_Run/internal/ui"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Initialise les constantes de la taille de la fenêtre
const (
	screenWidth     = 1024
	screenHeight    = 600
	ratioArrondiRec = float32(0.5)
	segmentRec      = int32(0)
)

var (
	fps        = 60
	toucheSaut = int32(rl.KeySpace) //Initialise la touche de saut

	frameCounter  = 0
	currentScreen = 0

	player1   = newPlayer()
	platform0 = newPlatform(rl.NewRectangle(screenWidth, screenHeight-150, 150, 50))
	platform1 = newPlatform(rl.NewRectangle(screenWidth+300, screenHeight-250, 150, 50))
	//platform2 = newPlatform(rl.NewRectangle(screenWidth/2, screenHeight-150, 150, 50))
	platforms = [2]*platform{platform0, platform1}

	isJumping    = false
	velocity     = float32(0.0)
	gravity      = float32(1.0)
	jumpStrength = float32(-20.0)
)

// Structure player avec sa hitbox et ses points de vie
type player struct {
	hitbox rl.Rectangle
	health int
}

type platform struct {
	hitbox rl.Rectangle
}

func newPlayer() *player {
	p := player{rl.NewRectangle(screenWidth/8, screenHeight-50.0, 50, 50), 3}
	return &p
}

func newPlatform(hitbox rl.Rectangle) *platform {
	p := platform{hitbox}
	return &p
}

func collision(player rl.Rectangle, platform rl.Rectangle) bool {
	return rl.CheckCollisionRecs(player, platform)
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "ABDEL RUN!!!")
	rl.SetTargetFPS(int32(fps)) //Initialise le nombre d'images par seconde
	rl.InitAudioDevice()        // Initialise le module audio

	// Charger l'image de fond et la musique une seule fois
	bgLogo := rl.LoadTexture("../assets/images/logo.jpg")
	bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
	bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
	playMusic := rl.LoadMusicStream("../assets/sounds/playSong.ogg")
	bgSettings := rl.LoadTexture("../assets/images/Settings.png")

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(bgMusic)
		rl.UpdateMusicStream(playMusic)

		switch currentScreen {
		case 0:
			currentScreen, frameCounter = ui.LogoScreen(currentScreen, frameCounter, bgLogo, bgMusic)
		case 1:
			currentScreen = ui.TitleScreen(currentScreen, ratioArrondiRec, segmentRec, bgImage, bgMusic)
		case 2:
			if !rl.IsMusicStreamPlaying(playMusic) {
				rl.UpdateMusicStream(playMusic)
				rl.PlayMusicStream(playMusic)
			}

			stringToucheSaut := strconv.FormatInt(int64(toucheSaut), 10)
			texteToucheSaut := "Press" + stringToucheSaut + " to jump"

			if collision(player1.hitbox, platforms[0].hitbox) {
				rl.DrawText("COLLISION!", 600, 250, 20, rl.Red)
			}
			//Saut du personnage
			if rl.IsKeyPressed(int32(toucheSaut)) && !isJumping {
				isJumping = true
				velocity = jumpStrength
			}

			//Retombée du personnage
			if isJumping {
				//frate := float32(velocity * (float32(fps) * rl.GetFrameTime()))

				player1.hitbox.Y += velocity
				velocity += gravity
				fmt.Println(player1.hitbox.Y)

				for i := 0; i < len(platforms); i++ {
					if collision(player1.hitbox, platforms[i].hitbox) && player1.hitbox.Y-player1.hitbox.Y < platforms[i].hitbox.Y-player1.hitbox.Height {
						velocity = platforms[i].hitbox.Height / 2
						player1.hitbox.Y += velocity
						velocity = 1.0
					}

					if player1.hitbox.Y > screenHeight-player1.hitbox.Height {
						player1.hitbox.Y = screenHeight - player1.hitbox.Height
						isJumping = false
					}
				}
				// if player.hitbox.Y >= player2.Y-player2.Height && player.hitbox.X+50 > player2.X && !collision {
				// 	velocity = 1.0
				// }
			}

			deplacementPlatform := float32(10)

			for i := 0; i < len(platforms); i++ {
				if collision(player1.hitbox, platforms[i].hitbox) {
					deplacementPlatform = 0

				}

				if player1.hitbox.Y <= platforms[i].hitbox.Y && collision(player1.hitbox, platforms[i].hitbox) {
					player1.hitbox.Y = platforms[i].hitbox.Y - platforms[i].hitbox.Height
					isJumping = false
				}
				platforms[i].hitbox.X -= deplacementPlatform
				if platforms[i].hitbox.X <= 0-platforms[i].hitbox.Width {
					platforms[i].hitbox.X = float32(rl.GetScreenWidth())
				}
			}

			// for i := 0; i < len(platforms); i++ {

			// }

			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawText(texteToucheSaut, 10, 0, 20, rl.Gray)
			rl.DrawRectangleRec(player1.hitbox, rl.Red)
			rl.DrawRectangleRec(platforms[0].hitbox, rl.Green)
			rl.DrawRectangleRec(platforms[1].hitbox, rl.Blue)

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
							player1.hitbox.Y = screenHeight - 50.0
							platform1.hitbox.X = screenWidth
							rl.StopMusicStream(playMusic)
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
				{rl.NewRectangle(screenWidth-220, screenHeight-360, 175, 40), "Change jump key"},
			}

			//Gestionnaire de FPS
			stringFPS := strconv.FormatInt(int64(fps), 10)
			rl.DrawText(stringFPS, screenWidth/2-150, screenHeight-420, 50, rl.Black)
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
							rl.UnloadMusicStream(bgMusic)
							rl.CloseAudioDevice()
							rl.CloseWindow()
							return
						case "Change jump key":
							currentScreen = 5
						}
					}
				}
				rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
				rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
			}
			rl.EndDrawing()
		case 4:
			rl.UnloadMusicStream(bgMusic)
			rl.CloseAudioDevice()
			rl.CloseWindow()
			return
		case 5:
			currentScreen, toucheSaut = ui.KeybindingScreen(currentScreen, ratioArrondiRec, segmentRec, bgLogo, bgMusic)
		}
	}
	rl.UnloadTexture(bgImage)
	rl.UnloadMusicStream(bgMusic)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
