package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

//raylib.Texture2D texture = LoadTexture("resources/raylib_logo.png");

func main() {
	raylib.InitWindow(screenWidth, screenHeight, "Go Geometry Dash")
	raylib.SetTargetFPS(60)
	frameCounter := 0
	currentScreen := 0

	player := raylib.Rectangle{
		X:      screenWidth / 8,
		Y:      screenHeight - 50.0,
		Width:  50,
		Height: 50,
	}

	isJumping := false
	velocity := float32(0.0)
	gravity := float32(1.0)
	jumpStrength := float32(-20.0)

	for !raylib.WindowShouldClose() {
		switch currentScreen {
		case 0:
			raylib.BeginDrawing()
			raylib.ClearBackground(raylib.Red)
			raylib.EndDrawing()
			frameCounter++
			if frameCounter > 120 {
				currentScreen++
			}
		case 1:

			if raylib.IsKeyPressed(raylib.KeySpace) && !isJumping {
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
			raylib.BeginDrawing()
			raylib.ClearBackground(raylib.White)
			raylib.DrawText("PRESS SPACE to PAUSE MOVEMENT", 10, 0, 20, raylib.Gray)
			raylib.DrawRectangleRec(player, raylib.Red)
			//raylib.DrawTexture(texture, screenWidth/2-texture.width/2, screenHeight/2-texture.height/2, WHITE)
			raylib.EndDrawing()
		}

	}

	raylib.CloseWindow()
}
