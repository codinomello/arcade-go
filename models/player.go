package models

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position rl.Vector2 // Posição do jogador
	Speed    rl.Vector2 // Velocidade do jogador
	Width    float32    // Largura do jogador
	Height   float32    // Altura do jogador
	OnGround bool       // Indica se o jogador está no chão
}
