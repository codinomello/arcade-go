package functions

import (
	"github.com/codinomello/arcade-go/core/logic"
	"github.com/codinomello/arcade-go/core/models"
)

// Função para atualizar a lógica do jogo
func UpdateGame(player *models.Player, platforms []models.Platform) {
	// Controles do jogador
	// controls.HandlePlayerControls(player)

	// Verificar colisões com as bordas da tela
	logic.CheckScreenBounds(player)

	// Verificar colisões com plataformas
	logic.CheckPlatformCollisions(player, platforms)
}
