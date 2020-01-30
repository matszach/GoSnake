package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	g "github.com/matszach/raylib/snake/game"
)

func main() {

	rl.InitWindow(g.Unit * g.GameWidth, g.Unit * g.GameHeight, "GoSnake")
	rl.SetTargetFPS(15)

	game := g.Game{}
	game.Init()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if !game.GameOver() {
			game.Logic()
		}
		game.Draw()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}