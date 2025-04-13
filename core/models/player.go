package models

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	// Atributos do jogador
	Position     rl.Vector2 // Posição do jogador
	Velocity     rl.Vector2 // Velocidade do jogador
	Acceleration float32    // Aceleração do jogador
	JumpForce    float32    // Força do pulo do jogador
	Gravity      float32    // Gravidade do jogador
	Size         rl.Vector2 // Tamanho do jogador
	Color        rl.Color   // Cor do jogador
	Lives        int        // Vidas do jogador
	Score        int        // Pontuação do jogador

	// Habilidades do jogador
	IsJumping       bool // Indica se o jogador está no chão
	OnGround        bool // Indica se o jogador está no chão
	CanDoubleJump   bool // Indica se o jogador pode dar um segundo pulo
	IsDoubleJumping bool // Indica se o jogador está dando um segundo pulo
	Invincible      bool // Indica se o jogador está invencível
	InvincibleTimer int  // Contador de tempo de invencibilidade
	HasDoubleJump   bool // Indica se o jogador tem o segundo pulo
	HasDash         bool // Indica se o jogador tem o dash
	IsDashing       bool // Indica se o jogador está usando o dash
	DashTimer       int  // Contador de tempo do dash
	DashDirection   int  // Direção do dash (-1 esquerda, 1 direita)
}
