package logic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlatform(x, y, w, h float32) Platform {
	return Platform{Rect: rl.NewRectangle(x, y, w, h)}
}

func (plat Platform) Draw() {
	rl.DrawRectangleRec(plat.Rect, rl.DarkGray)
}
