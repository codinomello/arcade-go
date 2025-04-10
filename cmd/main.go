package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Plataformer 2D")
	icon := rl.LoadImage("../images/icons/gopher-icon.png")
	defer rl.UnloadImage(icon)
	rl.SetWindowIcon(*icon)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Atualização
		// Desenho
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Olá, Plataforma!", 10, 10, 20, rl.Black)

		rl.EndDrawing()
	}
}
