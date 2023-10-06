package ui

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
)

var (
	scoreStr1 = "Score : "
)

func GameoverScreen(currentScreen int, score int, ratioArrondiRec float32, segmentRec int32, bgImage rl.Texture2D, gameoverMusic rl.Music) (int, int) {
	rl.BeginDrawing()
	rl.DrawTexture(bgImage, 0, 0, rl.White)

	rl.UpdateMusicStream(gameoverMusic)
	rl.PlayMusicStream(gameoverMusic)

	scoreStr2 := strconv.Itoa(score)
	scoreStrTot := scoreStr1 + scoreStr2

	buttons := []struct {
		Bounds rl.Rectangle
		Text   string
	}{
		{rl.NewRectangle(screenWidth/2-75, screenHeight/2-40, 150, 40), "Restart"},
		{rl.NewRectangle(screenWidth/2-75, screenHeight/2+10, 150, 40), "Main menu"},
		{rl.NewRectangle(screenWidth/2-75, screenHeight/2+60, 150, 40), "Quit"},
	}

	for _, button := range buttons {
		color := rl.Yellow
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.Bounds) {
			color = rl.DarkGray
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				switch button.Text {
				case "Restart":
					rl.StopMusicStream(gameoverMusic)
					score = 0
					currentScreen = 2
				case "Main menu":
					score = 0
					currentScreen = 1
				case "Quit":
					currentScreen = 4
				}
			}
		}
		rl.DrawRectangleRounded(button.Bounds, ratioArrondiRec, segmentRec, color)
		rl.DrawText(button.Text, int32(button.Bounds.X+button.Bounds.Width/2)-rl.MeasureText(button.Text, 20)/2, int32(button.Bounds.Y+10), 20, rl.Black)
		rl.DrawText(scoreStrTot, 420, screenHeight/3, 40, rl.Black)
	}
	rl.EndDrawing()
	return currentScreen, score
}
