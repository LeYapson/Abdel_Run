package main

import (
	"fmt"
	"strconv"

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

	player1 := rl.Rectangle{
		X:      screenWidth / 8,
		Y:      screenHeight - 50.0,
		Width:  50,
		Height: 50,
	}

	player2 := rl.Rectangle{
		X:      screenWidth / 8,
		Y:      screenHeight - 200,
		Width:  75,
		Height: 50,
	}

	isJumping := false
	velocite := float32(0.0)
	gravite := float32(1.0)
	forceSaut := float32(-20.0)

	toucheSaut := rl.KeySpace

	for !rl.WindowShouldClose() {
		collision := rl.CheckCollisionRecs(player1, player2)
		if collision {
			rl.DrawText("COLLISION!", 600, 250, 20, rl.Red)
		}
		//Saut du personnage
		if rl.IsKeyPressed(int32(toucheSaut)) && !isJumping {
			isJumping = true
			velocite = forceSaut
		}

		//Retombée du personnage
		if isJumping {
			//frate := float32(velocity * (float32(fps) * rl.GetFrameTime()))
			if collision {
				velocite = player2.Height / 2
				player1.Y += velocite
				velocite = 1.0
			}
			player1.Y += velocite
			velocite += gravite
			fmt.Println(player1.Y)

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

		if player1.Y <= player2.Y && collision {
			player1.Y = player2.Y - player2.Height - 1
			isJumping = false
			// } else if player1.X <= player2.X && !collision {
			// 	player1.Y += velocite
			// 	velocite += gravite
			//
		} else {
			deplacementPlayer2 := float32(2.5)
			if collision {
				deplacementPlayer2 = 0
			}
			player2.X += deplacementPlayer2
			if player2.X >= float32(rl.GetScreenWidth()-700) {
				player2.X = 0
			}
		}

		//fmt.Println(rl.GetFrameTime())

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText(texteToucheSaut, 10, 0, 20, rl.Gray)
		rl.DrawRectangleRec(player1, rl.Red)
		rl.DrawRectangleRec(player2, rl.Blue)

		rl.EndDrawing()
	}
}
