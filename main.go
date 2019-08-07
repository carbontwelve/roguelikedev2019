package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"raylibtinkering/ui"
)

var Version = "v0.1"

func main() {
	// Setup
	//----------------------------------------------------------------------------------
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "/r/roguelikedev 2019")
	rl.SetTargetFPS(60)

	ui.LinkColours()

	// NOTE: Textures and Sounds MUST be loaded after Window/Audio initialization
	game := newEngine()

	// Main Loop
	//----------------------------------------------------------------------------------
	for !rl.WindowShouldClose() {
		frameTime := rl.GetFrameTime()
		state := game.PeekState()

		// Update
		//----------------------------------------------------------------------------------

		//if rl.IsKeyPressed(rl.KeyF) {
		//	rl.ToggleFullscreen()
		//}

		state.Update(frameTime)

		// Draw
		//----------------------------------------------------------------------------------

		rl.BeginDrawing()
		state.Draw(frameTime)

		rl.DrawText(fmt.Sprintf("Delta: %f", frameTime), 20, 20, 10, ui.ColourFg)
		rl.EndDrawing()

		if state.ShouldQuit() {
			break
		}
	}

	// Free resources & Exit
	//----------------------------------------------------------------------------------

	game.Unload()
	rl.CloseWindow()
	os.Exit(0)
}
