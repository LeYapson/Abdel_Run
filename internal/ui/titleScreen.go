package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

func TitleScreen(currentScreen int, ratioArrondiRec float32, segmentRec int32, bgImage rl.Texture2D, bgMusic rl.Music) int {
	rl.BeginDrawing()
	rl.DrawTexture(bgImage, 0, 0, rl.White)

	if !rl.IsMusicStreamPlaying(bgMusic) {
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
				case "Play":
					rl.StopMusicStream(bgMusic)
					currentScreen = 2
				case "Settings":
					currentScreen = 3
				case "Quit":
					currentScreen = 4
					//return
				}
			}
		}
		rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
		rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
	}
	rl.EndDrawing()
	return currentScreen
}
