package main

import (
	"github.com/codinomello/arcade-go/config"
	"github.com/codinomello/arcade-go/logic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Inicialização da janela
	window := config.LoadWindowConfig()
	rl.InitWindow(window.ScreenWidth, window.ScreenHeight, "Gopher, Go!")
	rl.SetTargetFPS(60)
	config.LoadWindowIcon("gopher", "png")

	// Inicialização do jogo
	player := logic.NewPlayer()
	platforms := []logic.Platform{
		logic.NewPlatform(0, 400, 800, 50),
		logic.NewPlatform(300, 300, 200, 20),
	}

	for !rl.WindowShouldClose() {
		// Atualização da lógica do jogo
		player.Update(platforms)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		player.Draw()
		for _, p := range platforms {
			p.Draw()
		}

		rl.EndDrawing()
	}
}
