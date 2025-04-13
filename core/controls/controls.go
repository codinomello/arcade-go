package controls

import (
	"github.com/codinomello/arcade-go/core/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player models.Player

func HandlePlayerControls(player *models.Player) {
	// Movimento horizontal
	if rl.IsKeyDown(rl.KeyRight) {
		player.Velocity.X = player.Acceleration
	} else if rl.IsKeyDown(rl.KeyLeft) {
		player.Velocity.X -= player.Acceleration
	} else {
		// Desaceleração horizontal quando nenhuma tecla é pressionada
		if player.Velocity.X > 0 {
			player.Velocity.X -= 0.5
			if player.Velocity.X < 0 {
				player.Velocity.X = 0
			}
		} else if player.Velocity.X < 0 {
			player.Velocity.X += 0.5
			if player.Velocity.X > 0 {
				player.Velocity.X = 0
			}
		}
	}
}

// func (player *Player) Update(dt float32, platforms []rl.Rectangle) {
// 	// Guarda o estado atual de contato com o solo
// 	player.WasOnGround = p.OnGround

// 	// Controle de timers
// 	p.UpdateTimers()

// 	// Se estiver dashando, a lógica é diferente
// 	if p.IsDashing {
// 		p.UpdateDash(deltaTime)
// 	} else {
// 		// Movimento horizontal com aceleração suave
// 		p.HandleHorizontalMovement(deltaTime)

// 		// Processa entrada de pulo e aplica coyote time
// 		p.HandleJump(deltaTime)

// 		// Tenta iniciar um dash
// 		p.TryDash()

// 		// Aplica gravidade se não estiver no chão
// 		if !p.OnGround {
// 			p.Velocity.Y += p.Gravity * deltaTime
// 		}
// 	}

// 	// Atualiza a posição com base na velocidade
// 	p.Position.X += p.Velocity.X * deltaTime
// 	p.Position.Y += p.Velocity.Y * deltaTime

// 	// Detecta e resolve colisões
// 	p.CheckCollisions(platforms)

// 	// Atualiza o estado de OnGround após verificar colisões
// 	// Se não há mais contato com o solo mas estava no solo no frame anterior,
// 	// ativa o coyote time
// 	if p.WasOnGround && !p.OnGround {
// 		p.CoyoteTimeCounter = COYOTE_TIME
// 	}
// }
