package ui

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Gameplay(currentScreen int, screenWidth float32, screenHeight float32, fps int) int {

	ratioArrondiRec := float32(0.5)
	segmentRec := int32(0)
	isJumping := false
	velocity := float32(0.0)
	gravity := float32(1.0)
	jumpStrength := float32(-20.0)
	toucheSaut := rl.KeySpace //Initialise la touche de saut

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
	//Saut du personnage
	if rl.IsKeyPressed(int32(toucheSaut)) && !isJumping {
		isJumping = true
		velocity = jumpStrength
		//fmt.Println(velocity)
	}

	//Retombée du personnage
	if isJumping {
		//fmt.Println(velocity)
		//frate := float32(velocity * (float32(fps) * rl.GetFrameTime()))
		player1.Y += velocity
		velocity += gravity
		if player1.Y > screenHeight-player1.Height {
			player1.Y = screenHeight - player1.Height
			isJumping = false
		}
	}
	stringToucheSaut := strconv.FormatInt(int64(toucheSaut), 10)
	texteToucheSaut := "Press" + stringToucheSaut + " to jump"

	fmt.Println(rl.GetFrameTime())

	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawText(texteToucheSaut, 10, 0, 20, rl.Gray)
	rl.DrawRectangleRec(player1, rl.Red)
	rl.DrawRectangleRec(player2, rl.Blue)

	collision := rl.CheckCollisionRecs(player1, player2)

	if collision {
		rl.DrawText("COLLISION!", 600, 250, 20, rl.Red)
	}

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
	return currentScreen

}
