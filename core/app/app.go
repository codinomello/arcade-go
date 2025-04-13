package app

import (
	"fmt"

	"github.com/codinomello/arcade-go/core/config"
	"github.com/codinomello/arcade-go/core/functions"
	"github.com/codinomello/arcade-go/core/logic"
	"github.com/codinomello/arcade-go/core/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Name string
	// TODO: outros componentes
}

// NewGame cria uma nova instância do jogo
func NewGame() *Game {
	// Inicialização da janela
	window := config.LoadWindowConfig()
	rl.InitWindow(window.ScreenWidth, window.ScreenHeight, window.Title)
	rl.SetTargetFPS(window.TargetFPS)
	config.LoadWindowIcon(window.Icon, window.IconFormat)

	return &Game{
		Name: window.Title,
	}
}

// Run inicia o loop principal do jogo
func (game *Game) Run() {
	// Inicialização do jogo
	player := logic.NewPlayer()
	platforms := []models.Platform{
		logic.NewPlatform(0, 400, 800, 50),
		logic.NewPlatform(300, 300, 200, 20),
	}

	for !rl.WindowShouldClose() {
		// dt := rl.GetFrameTime()

		// Atualização da lógica do jogo
		functions.UpdateGame(&player, platforms)

		// Desenho
		rl.BeginDrawing()

		// Limpar a tela
		rl.ClearBackground(rl.RayWhite)

		// Desenhar player
		logic.DrawPlayer(player)

		// Desenhar plataformas
		for _, platform := range platforms {
			logic.DrawPlatform(platform)
		}

		// Informações de debug
		rl.DrawText(fmt.Sprintf("Posição: (%.1f, %.1f)", player.Position.X, player.Position.Y), 10, 10, 20, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("Velocidade: (%.1f, %.1f)", player.Velocity.X, player.Velocity.Y), 10, 40, 20, rl.DarkGray)
		rl.DrawText("Controles: Setas para mover, Espaço para pular", 10, 70, 20, rl.DarkGray)

		// Encerra o desenho
		rl.EndDrawing()
	}
}

// Shutdown finaliza o jogo e libera os recursos utilizados
func (game *Game) Shutdown() {
	// Limpeza de recursos
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
