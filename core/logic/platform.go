package logic

import (
	"github.com/codinomello/arcade-go/core/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlatform(x, y, w, h float32) models.Platform {
	return models.Platform{Rectangle: rl.NewRectangle(x, y, w, h)}
}

func DrawPlatform(platform models.Platform) {
	rl.DrawRectangleRec(platform.Rectangle, rl.DarkGray)
}
