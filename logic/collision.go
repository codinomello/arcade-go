package logic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) CheckCollision(platforms []Platform) {
	p.OnGround = false
	playerRect := rl.NewRectangle(p.Position.X, p.Position.Y, p.Width, p.Height)

	for _, platform := range platforms {
		if rl.CheckCollisionRecs(playerRect, platform.Rect) {
			p.Position.Y = platform.Rect.Y - p.Height
			p.Speed.Y = 0
			p.OnGround = true
		}
	}
}
