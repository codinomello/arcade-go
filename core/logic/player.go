package logic

import (
	"github.com/codinomello/arcade-go/core/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Gravity   = 0.5
	JumpForce = -10
	MoveSpeed = 5
)

func NewPlayer() models.Player {
	return models.Player{
		Position:     rl.Vector2{X: 100, Y: 100}, // Posição inicial do jogador
		Velocity:     rl.NewVector2(0, 0),        // Velocidade do jogador
		Acceleration: 4.0,                        // Aceleração do jogador
		JumpForce:    -12.0,                      // Força do pulo do jogador
		Gravity:      0.5,                        // Gravidade do jogador
		Lives:        3,                          // Vidas do jogador
		Size:         rl.NewVector2(50, 50),      // Tamanho do jogador
		Color:        rl.SkyBlue,                 // Cor do jogador
		IsJumping:    false,                      // Inicialmente, o jogador não está no chão
		Score:        0,                          // Pontuação do jogador

	}
}

func DrawPlayer(player models.Player) {
	rl.DrawRectangle(int32(player.Position.X), int32(player.Position.Y), int32(player.Size.X), int32(player.Size.Y), rl.Blue)
}
