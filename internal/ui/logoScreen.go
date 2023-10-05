package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LogoScreen(currentScreen int, frameCounter int, bgLogo rl.Texture2D, bgMusic rl.Music) (int, int) {
	rl.BeginDrawing()
	rl.DrawTexture(bgLogo, 0, 0, rl.White)
	rl.EndDrawing()
	frameCounter++
	if frameCounter > 120 {
		currentScreen++
		rl.PlayMusicStream(bgMusic)
	}
	return currentScreen, frameCounter
}
