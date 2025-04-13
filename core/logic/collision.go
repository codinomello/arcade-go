package logic

import (
	"math"

	"github.com/codinomello/arcade-go/core/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckPlatformCollisions(player *models.Player, platforms []models.Platform) {
	playerRect := rl.NewRectangle(player.Position.X, player.Position.Y, player.Size.X, player.Size.Y)

	for _, platform := range platforms {
		// Verificar se há colisão com a plataforma
		if rl.CheckCollisionRecs(playerRect, platform.Rectangle) {
			// Calcular as bordas do jogador e da plataforma
			playerBottom := player.Position.Y + player.Size.Y
			playerTop := player.Position.Y
			playerRight := player.Position.X + player.Size.X
			playerLeft := player.Position.X

			platformBottom := platform.Rectangle.Y + platform.Rectangle.Height
			platformTop := platform.Rectangle.Y
			platformRight := platform.Rectangle.X + platform.Rectangle.Width
			platformLeft := platform.Rectangle.X

			// Calcular penetração em cada direção
			bottomCollision := platformTop - playerBottom
			topCollision := platformBottom - playerTop
			leftCollision := platformRight - playerLeft
			rightCollision := platformLeft - playerRight

			// Encontrar a menor penetração para determinar a direção da colisão
			minPenetration := bottomCollision
			collisionDirection := "bottom"

			if math.Abs(float64(topCollision)) < math.Abs(float64(minPenetration)) {
				minPenetration = topCollision
				collisionDirection = "top"
			}

			if math.Abs(float64(leftCollision)) < math.Abs(float64(minPenetration)) {
				minPenetration = leftCollision
				collisionDirection = "left"
			}

			if math.Abs(float64(rightCollision)) < math.Abs(float64(minPenetration)) {
				minPenetration = rightCollision
				collisionDirection = "right"
			}

			// Resolver a colisão com base na direção
			switch collisionDirection {
			case "bottom":
				// Colisão com a parte superior da plataforma (mais comum)
				player.Position.Y = platformTop - player.Size.Y
				player.Velocity.Y = 0
				player.IsJumping = false
			case "top":
				// Colisão com a parte inferior da plataforma
				player.Position.Y = platformBottom
				player.Velocity.Y = 0
			case "left":
				// Colisão com o lado direito da plataforma
				player.Position.X = platformRight
				player.Velocity.X = 0
			case "right":
				// Colisão com o lado esquerdo da plataforma
				player.Position.X = platformLeft - player.Size.X
				player.Velocity.X = 0
			}
		}
	}
}
