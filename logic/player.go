package logic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Gravity   = 0.5
	JumpForce = -10
	MoveSpeed = 5
)

func NewPlayer() Player {
	return Player{
		Position: rl.Vector2{X: 100, Y: 100}, // Posição inicial do jogador
		Speed:    rl.NewVector2(0, 0),        // Velocidade do jogador
		Width:    50,                         // Largura do jogador
		Height:   50,                         // Altura do jogador
		OnGround: false,                      // Inicialmente, o jogador não está no chão
	}
}

func (p *Player) Update(platforms []Platform) {
	if rl.IsKeyDown(rl.KeyRight) {
		p.Position.X += MoveSpeed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Position.X -= MoveSpeed
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.OnGround {
		p.Speed.Y = JumpForce
		p.OnGround = false
	}

	p.Speed.Y += Gravity
	p.Position.Y += p.Speed.Y

	p.CheckCollision(platforms)
}

func (p Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), int32(p.Width), int32(p.Height), rl.Blue)
}
