package logic

import (
	"github.com/codinomello/arcade-go/core/config"
	"github.com/codinomello/arcade-go/core/models"
)

// Função para verificar as bordas da tela
func CheckScreenBounds(player *models.Player) {
	window := config.LoadWindowConfig()

	// Borda esquerda
	if player.Position.X < 0 {
		player.Position.X = 0
		player.Velocity.X = 0
	}

	// Borda direita
	if player.Position.X+player.Size.X > float32(window.ScreenWidth) {
		player.Position.X = float32(window.ScreenWidth) - player.Size.X
		player.Velocity.X = 0
	}

	// Borda inferior (opcional, normalmente coberta pelas plataformas)
	if player.Position.Y+player.Size.Y > float32(window.ScreenHeight) {
		player.Position.Y = float32(window.ScreenHeight) - player.Size.Y
		player.Velocity.Y = 0
		player.IsJumping = false
	}
}
