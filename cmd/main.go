package main

import (
	"strconv"

	ui "github.com/LeYapson/Abdel_Run/internal/ui"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

var fps = 60

func main() {
	rl.InitWindow(screenWidth, screenHeight, "ABDEL RUN!!!")
	rl.SetTargetFPS(int32(fps))
	rl.InitAudioDevice() // Initialise le module audio

	frameCounter := 0
	currentScreen := 0

	player1 := rl.Rectangle{
		X:      screenWidth / 8,
		Y:      screenHeight - 50.0,
		Width:  50,
		Height: 50,
	}

	player2 := rl.Rectangle{
		X:      screenWidth / 8,
		Y:      screenHeight - 150,
		Width:  50,
		Height: 50,
	}

	isJumping := false
	velocity := float32(0.0)
	gravity := float32(1.0)
	jumpStrength := float32(-20.0)

	ratioArrondiRec := float32(0.5)
	segmentRec := int32(0)

	toucheSaut := rl.KeySpace //Initialise la touche de saut
	//animFrames := int32(0)
	//p := &animFrames

	// Charger l'image de fond et la musique une seule fois
	bgLogo := rl.LoadTexture("../assets/images/logo.jpg")
	bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
	bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
	playMusic := rl.LoadMusicStream("../assets/sounds/playSong.ogg")
	bgSettings := rl.LoadTexture("../assets/images/Settings.png")
	//gifAbdel := rl.LoadImageAnim("../assets/abdel_run.gif", p)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(bgMusic)
		rl.UpdateMusicStream(playMusic)

		switch currentScreen {
		case 0:
			currentScreen, frameCounter = ui.LogoScreen(currentScreen, frameCounter, bgLogo, bgMusic)
		case 1:

			currentScreen = ui.TitleScreen(currentScreen, bgImage, bgMusic)

		case 2:
			if !rl.IsMusicStreamPlaying(playMusic) {
				rl.UpdateMusicStream(playMusic)
				rl.PlayMusicStream(playMusic)
			}
			collision := rl.CheckCollisionRecs(player1, player2)

			if collision {
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
				player1.Y += velocity
				velocity += gravity
				if player1.Y > screenHeight-player1.Height {
					player1.Y = screenHeight - player1.Height
					isJumping = false
					// } else if (collision){
					// 	player1.Y =
					// }
				}
			}
			stringToucheSaut := strconv.FormatInt(int64(toucheSaut), 10)
			texteToucheSaut := "Press" + stringToucheSaut + " to jump"

			//fmt.Println(rl.GetFrameTime())

			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawText(texteToucheSaut, 10, 0, 20, rl.Gray)
			rl.DrawRectangleRec(player1, rl.Red)
			rl.DrawRectangleRec(player2, rl.Blue)

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

			//currentScreen = lgs.Gameplay(currentScreen, screenWidth, screenHeight, fps)
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
		case 4:
			rl.UnloadMusicStream(bgMusic)
			rl.CloseAudioDevice()
			rl.CloseWindow()
		}
	}
	rl.UnloadTexture(bgImage)
	rl.UnloadMusicStream(bgMusic)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
