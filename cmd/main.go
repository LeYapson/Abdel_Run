package main

import (
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

	dimPlayer1 = rl.NewRectangle(screenWidth/8, screenHeight-50.0, 50, 50)
	player1    = newPlayer(dimPlayer1)

	dimPlatform0 = rl.NewRectangle(screenWidth, screenHeight-150, 90, 50)
	dimPlatform1 = rl.NewRectangle(screenWidth+200, screenHeight-250, 200, 50)
	dimPlatform2 = rl.NewRectangle(screenWidth+450, screenHeight-350, 140, 50)
	dimPlatform3 = rl.NewRectangle(screenWidth+600, screenHeight-200, 140, 50)
	platform0    = newPlatform(dimPlatform0)
	platform1    = newPlatform(dimPlatform1)
	platform2    = newPlatform(dimPlatform2)
	platform3    = newPlatform(dimPlatform3)
	platforms    = [4]*platform{platform0, platform1, platform2, platform3}

	dimEnnemi0 = rl.NewRectangle(screenWidth, screenHeight-50, 30, 30)
	ennemi0    = newEnnemi(dimEnnemi0)
	ennemis    = [1]*ennemi{ennemi0}

	isJumping     = false
	velocity      = float32(0.0)
	gravity       = float32(1.0)
	jumpStrength  = float32(-20.0)
	score         = 0
	deathCount    = 0
	deathCountStr string

	platformsAndEnnemisSpeed     = float32(10)
	platformsAndEnnemisSpeedAcce = platformsAndEnnemisSpeed
)

// Structure player avec sa hitbox et ses points de vie
type player struct {
	hitbox rl.Rectangle
	health int
}

type platform struct {
	hitbox rl.Rectangle
}

type ennemi struct {
	hitbox rl.Rectangle
}

func newPlayer(hitbox rl.Rectangle) *player {
	p := player{hitbox, 1}
	return &p
}

func newEnnemi(hitbox rl.Rectangle) *ennemi {
	e := ennemi{hitbox}
	return &e
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

	// Charger les images et les musiques une seule fois
	bgLogo := rl.LoadTexture("../assets/images/logo.jpg")
	bgImage := rl.LoadTexture("../assets/images/TitleScreen.jpg")
	bgSettings := rl.LoadTexture("../assets/images/Settings.png")
	bgMapKey := rl.LoadTexture("../assets/images/bgMapKey.jpg")
	bgPlay := rl.LoadTexture("../assets/images/bgPlay.png")
	bgtoucheW := rl.LoadTexture("../assets/images/bgtoucheW.jpg")
	bgMusic := rl.LoadMusicStream("../assets/sounds/AbdelRunSong.ogg")
	bgGameOver := rl.LoadTexture("../assets/images/GameOver.jpg")
	playMusic := rl.LoadMusicStream("../assets/sounds/playSong.ogg")
	playW := rl.LoadMusicStream("../assets/sounds/toucheW.ogg")
	gameoverMusic := rl.LoadMusicStream("../assets/sounds/gameoverMusic.ogg")

	//S'exécute à chaque frame
	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(bgMusic)
		rl.UpdateMusicStream(playMusic)

		switch currentScreen {
		//case 0 = LOADING GAME SCREEN
		case 0:
			currentScreen, frameCounter = ui.LogoScreen(currentScreen, frameCounter, bgLogo, bgMusic)
		//case 1 = MAIN MENU SCREEN
		case 1:
			currentScreen = ui.TitleScreen(currentScreen, ratioArrondiRec, segmentRec, bgImage, bgMusic)
		//case 2 = PLAY SCREEN
		case 2:
			if !rl.IsMusicStreamPlaying(playMusic) {
				rl.UpdateMusicStream(playMusic)
				rl.PlayMusicStream(playMusic)
			}

			if rl.IsKeyDown(90) {
				rl.SetTargetFPS(144)
				rl.BeginDrawing()
				rl.DrawTexture(bgtoucheW, 0, 0, rl.White)
				rl.StopMusicStream(playMusic)
				rl.UpdateMusicStream(playW)
				rl.PlayMusicStream(playW)

			} else {
				rl.SetTargetFPS(int32(fps))
				rl.BeginDrawing()
				rl.DrawTexture(bgPlay, 0, 0, rl.White)
			}

			score += 1
			if score%200 == 0 {
				platformsAndEnnemisSpeedAcce++
			}
			scoreStr := strconv.Itoa(score)

			stringToucheSaut := strconv.FormatInt(int64(toucheSaut), 10)
			texteToucheSaut := "Press " + stringToucheSaut + " to jump"
			textToucheW := "Press W for using the TIME POWER"
			rl.DrawText(texteToucheSaut, 10, 0, 20, rl.Yellow)
			rl.DrawText(textToucheW, 650, 0, 20, rl.Yellow)

			//Saut du personnage
			if rl.IsKeyPressed(int32(toucheSaut)) && !isJumping {
				isJumping = true
				velocity = jumpStrength

			}

			//Gestion des déplacements et des collisions avec les plateformes
			for i := 0; i < len(platforms); i++ {
				if collision(player1.hitbox, platforms[i].hitbox) {
					platformsAndEnnemisSpeed = 0
				}
				platformsAndEnnemisSpeed = platformsAndEnnemisSpeedAcce
				//Gestion
				if player1.hitbox.Y <= platforms[i].hitbox.Y && collision(player1.hitbox, platforms[i].hitbox) {
					player1.hitbox.Y = platforms[i].hitbox.Y - platforms[i].hitbox.Height
					isJumping = false

				}
				//Gestion réapparition de la plateforme quand elle sort de l'écran
				platforms[i].hitbox.X -= platformsAndEnnemisSpeed
				if platforms[i].hitbox.X <= 0-platforms[i].hitbox.Width {
					platforms[i].hitbox.X = float32(rl.GetScreenWidth())
				}

			}

			//Retombée du personnage
			if isJumping {
				player1.hitbox.Y += velocity
				velocity += gravity
				for i := 0; i < len(platforms); i++ {
					if collision(player1.hitbox, platforms[i].hitbox) {
						velocity = platforms[i].hitbox.Height / 2
						player1.hitbox.Y += velocity
						velocity = 1.0
					}

					//Maintien du joueur en bas de l'écran
					if player1.hitbox.Y > screenHeight-player1.hitbox.Height {
						player1.hitbox.Y = screenHeight - player1.hitbox.Height
						isJumping = false

					}
				}
			}

			//Gestions des déplacements et des collisions avec les ennemis
			for i := 0; i < len(ennemis); i++ {
				if collision(player1.hitbox, ennemis[i].hitbox) {
					player1.health--
				}
				ennemis[i].hitbox.X -= platformsAndEnnemisSpeed
				if ennemis[i].hitbox.X <= 0-ennemis[i].hitbox.Width {
					ennemis[i].hitbox.X = float32(rl.GetScreenWidth())
				}
			}

			if player1.health == 0 {
				deathCountStr = "Death count : "
				deathCount++
				deathCountStr += strconv.Itoa(deathCount)
				rl.StopMusicStream(playMusic)
				currentScreen = 6
			}

			rl.DrawRectangleRec(player1.hitbox, rl.Blue)

			//Dessiner toutes les plateformes
			for i := 0; i < len(platforms); i++ {
				rl.DrawRectangleRec(platforms[i].hitbox, rl.Green)
			}

			rl.DrawRectangleRec(ennemis[0].hitbox, rl.Red)
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
							score = 0
							player1 = newPlayer(rl.NewRectangle(screenWidth/8, screenHeight-50.0, 50, 50))

							platform0 = newPlatform(dimPlatform0)
							platform1 = newPlatform(dimPlatform1)
							platform2 = newPlatform(dimPlatform2)
							platform3 = newPlatform(dimPlatform3)
							platforms = [4]*platform{platform0, platform1, platform2, platform3}

							ennemi0 = newEnnemi(rl.NewRectangle(screenWidth, screenHeight-50, 30, 30))
							ennemis = [1]*ennemi{ennemi0}
							rl.StopMusicStream(playMusic)
							currentScreen = 1
						}
					}
				}
				//Affichage du bouton
				rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
				rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
				rl.DrawText(scoreStr, screenWidth-200, 50, 50, rl.White)
				rl.DrawText(deathCountStr, screenWidth-200, 100, 25, rl.White)
			}
			rl.EndDrawing()
		//case 3 = SETTINGS SCREEN
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
		//case 4 = QUIT THE GAME
		case 4:
			rl.UnloadMusicStream(bgMusic)
			rl.CloseAudioDevice()
			rl.CloseWindow()
			return
		//case 5 = Settings ---> key mapping SCREEN
		case 5:
			rl.BeginDrawing()
			rl.DrawTexture(bgMapKey, 0, 0, rl.White)
			currentScreen, toucheSaut = ui.KeybindingScreen(currentScreen, ratioArrondiRec, segmentRec, bgLogo, bgMusic)

		//case 6 = GAME OVER SCREEN
		case 6:
			player1 = newPlayer(dimPlayer1)

			platform0 = newPlatform(dimPlatform0)
			platform1 = newPlatform(dimPlatform1)
			platform2 = newPlatform(dimPlatform2)
			platform3 = newPlatform(dimPlatform3)
			platforms = [4]*platform{platform0, platform1, platform2, platform3}

			ennemi0 = newEnnemi(dimEnnemi0)
			ennemis = [1]*ennemi{ennemi0}
			currentScreen, score = ui.GameoverScreen(currentScreen, score, ratioArrondiRec, segmentRec, bgGameOver, gameoverMusic)
		}

	}
	rl.UnloadTexture(bgImage)
	rl.UnloadMusicStream(bgMusic)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
