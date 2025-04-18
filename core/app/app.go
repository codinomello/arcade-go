package app

import (
	"fmt"

	"github.com/codinomello/arcade-go/core/config"
	"github.com/codinomello/arcade-go/core/entities"
	"github.com/codinomello/arcade-go/core/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Name  string
	Level string
	Score int
	// TODO: outros componentes
}

// Cria uma nova instância do jogo
func NewGame() *Game {
	// Inicialização da janela
	window := config.LoadWindowConfig()
	rl.InitWindow(int32(window.ScreenWidth), int32(window.ScreenHeight), window.Title)
	rl.SetTargetFPS(window.TargetFPS)
	rl.SetExitKey(rl.KeyEscape)

	// Carregar o ícone da janela
	config.LoadWindowIcon(window.Icon, window.IconFormat)

	return &Game{
		Name:  window.Title,
		Level: "level-0",
		Score: 0,
	}
}

// Inicia o loop principal do jogo
func (game *Game) Run() {
	// Inicialização do jogo
	player := entities.NewPlayer()
	platforms := []objects.Platform{
		objects.NewPlatform(0, 400, 800, 50),
		objects.NewPlatform(300, 300, 200, 20),
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// Atualização do jogador
		player.Update(dt, platforms)

		// Desenho
		rl.BeginDrawing()

		// Limpar a tela
		rl.ClearBackground(rl.RayWhite)

		// Desenhar player
		player.Draw()

		// Desenhar plataformas
		for _, platform := range platforms {
			platform.Draw()
		}

		// Informações de debug
		rl.DrawText(fmt.Sprintf("Posição: (%.1f, %.1f)", player.Position.X, player.Position.Y), 10, 10, 20, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("Velocidade: (%.1f, %.1f)", player.Velocity.X, player.Velocity.Y), 10, 40, 20, rl.DarkGray)
		rl.DrawText("Controles: Setas para mover, Espaço para pular", 10, 70, 20, rl.DarkGray)

		// Encerra o desenho
		rl.EndDrawing()
	}
}

// Finaliza o jogo e libera os recursos utilizados
func (game *Game) Shutdown() {
	// Limpeza de recursos
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
