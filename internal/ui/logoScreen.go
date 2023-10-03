package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LogoScreen(frameCounter int, currentScreen int, bgMusic rl.Music, bgLogo rl.Texture2D) (int, int) {
	rl.BeginDrawing()
	rl.DrawTexture(bgLogo, 0, 0, rl.White)
	rl.EndDrawing()
	frameCounter++
	if frameCounter > 120 {
		currentScreen++
		rl.PlayMusicStream(bgMusic)
	}
	return frameCounter, currentScreen

}
