package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Platform struct {
	// Atributos da plataforma
	Rectangle rl.Rectangle // Retângulo que representa a plataforma
	Position  rl.Vector2   // Posição da plataforma
	Size      rl.Vector2   // Tamanho da plataforma
	Color     rl.Color     // Cor da plataforma

	// Atributos de movimento da plataforma
	Moving        bool       // Indica se a plataforma está em movimento
	Velocity      rl.Vector2 // Velocidade da plataforma
	StartPosition rl.Vector2 // Posição inicial do movimento da plataforma
	EndPosition   rl.Vector2 // Posição inicial e final do movimento da plataforma
	Progress      float32    // Progresso do movimento da plataforma
}

// Cria uma nova plataforma
func NewPlatform(x, y, w, h float32) Platform {
	return Platform{Rectangle: rl.NewRectangle(x, y, w, h)}
}

// Desenha a plataforma na tela
func (platform *Platform) Draw() {
	rl.DrawRectangleRec(platform.Rectangle, rl.DarkGray)
}
