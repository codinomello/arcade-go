package entities

import (
	"math"

	"github.com/codinomello/arcade-go/core/constants"
	"github.com/codinomello/arcade-go/core/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	// Atributos básicos
	Position     rl.Vector2 // Posição do jogador
	Velocity     rl.Vector2 // Velocidade do jogador
	Acceleration float32    // Aceleração do jogador
	JumpForce    float32    // Força do pulo do jogador
	Gravity      float32    // Gravidade aplicada ao jogador
	Size         rl.Vector2 // Tamanho do jogador
	Color        rl.Color   // Cor do jogador
	Lives        int        // Quantidade de vidas
	Score        int        // Pontuação atual

	// Estados e habilidades
	IsJumping       bool // Se está pulando
	OnGround        bool // Se está no chão
	WasOnGround     bool // Estado anterior (para coyote time)
	CanDoubleJump   bool // Se pode dar um duplo pulo
	IsDoubleJumping bool // Se está no segundo pulo
	Invincible      bool // Estado de invencibilidade
	InvincibleTimer int  // Timer de invencibilidade
	HasDoubleJump   bool // Se possui habilidade de duplo pulo
	HasDash         bool // Se possui habilidade de dash
	IsDashing       bool // Se está usando dash
	DashTimer       int  // Duração do dash
	DashDirection   int  // Direção do dash (-1: esq, 1: dir)

	// Contadores e timers
	DashCooldown      int // Tempo de recarga do dash
	JumpBufferCounter int // Buffer de entrada de pulo
	CoyoteTimeCounter int // Tempo para coyote time
	FacingDirection   int // Direção que o jogador está virado
}

const (
	MAX_SPEED         = 8.0
	DECELERATION      = 0.8
	DASH_DURATION     = 15
	DASH_SPEED        = 20.0
	DASH_COOLDOWN     = 30
	INVINCIBLE_TIME   = 120
	JUMP_BUFFER_TIME  = 10
	COYOTE_TIME       = 7
	DOUBLE_JUMP_FORCE = 10.0
	SCREEN_WIDTH      = 800
	SCREEN_HEIGHT     = 600
)

func NewPlayer() Player {
	return Player{
		Position:        rl.Vector2{X: 100, Y: 100}, // Posição inicial do jogador
		Velocity:        rl.NewVector2(0, 0),        // Velocidade inicial zero
		Acceleration:    4.0,                        // Aceleração padrão
		JumpForce:       -12.0,                      // Força do pulo (negativo para ir para cima)
		Gravity:         0.5,                        // Gravidade aplicada
		Size:            rl.NewVector2(50, 50),      // Tamanho do retângulo
		Color:           rl.SkyBlue,                 // Cor do jogador
		Lives:           3,                          // Vidas iniciais
		Score:           0,                          // Pontuação inicial
		IsJumping:       false,                      // Começa sem estar pulando
		OnGround:        false,                      // Começa no ar
		WasOnGround:     false,                      // Começa sem estar no chão
		CanDoubleJump:   false,                      // Não pode dar duplo pulo inicialmente
		IsDoubleJumping: false,                      // Não está no segundo pulo
		Invincible:      false,                      // Não é invencível inicialmente

	}
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), int32(p.Size.X), int32(p.Size.Y), p.Color)
}

// Atualiza a lógica do jogador a cada frame
func (p *Player) Update(dt float32, platforms []objects.Platform) {
	// Guarda o estado anterior de contato com o chão
	p.WasOnGround = p.OnGround

	// Atualiza todos os timers
	p.UpdateTimers()

	if p.IsDashing {
		p.UpdateDash(dt)
	} else {
		// Movimento horizontal normal
		p.HandleHorizontalMovement(dt)

		// Lógica de pulo
		p.HandleJump(dt)

		// Tenta iniciar dash
		p.TryDash()

		// Aplica gravidade se não estiver no chão
		if !p.OnGround {
			p.Velocity.Y += p.Gravity * dt
		}
	}

	// Atualiza posição baseada na velocidade
	p.Position.X += p.Velocity.X * dt
	p.Position.Y += p.Velocity.Y * dt

	// Verifica colisões
	p.CheckCollisions(platforms)
}

// UpdateTimers gerencia todos os contadores de tempo do jogador
func (p *Player) UpdateTimers() {
	if p.Invincible {
		p.InvincibleTimer--
		if p.InvincibleTimer <= 0 {
			p.Invincible = false
		}
	}

	if p.DashCooldown > 0 {
		p.DashCooldown--
	}

	if p.JumpBufferCounter > 0 {
		p.JumpBufferCounter--
	}

	if p.CoyoteTimeCounter > 0 {
		p.CoyoteTimeCounter--
	}
}

// HandleHorizontalMovement processa o movimento horizontal
func (p *Player) HandleHorizontalMovement(dt float32) {
	movingRight := rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)
	movingLeft := rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)

	// Controle reduzido no ar
	controlMultiplier := float32(1.0)
	if !p.OnGround {
		controlMultiplier = constants.AIR_RESISTANCE
	}

	var accelFactor float32 = 0
	if movingRight {
		accelFactor = 1
		p.FacingDirection = 1
	} else if movingLeft {
		accelFactor = -1
		p.FacingDirection = -1
	}

	// Aplica aceleração ou desaceleração
	if accelFactor != 0 {
		p.Velocity.X += accelFactor * p.Acceleration * dt * controlMultiplier
	} else {
		if p.Velocity.X > 0 {
			p.Velocity.X -= DECELERATION * dt
			if p.Velocity.X < 0 {
				p.Velocity.X = 0
			}
		} else if p.Velocity.X < 0 {
			p.Velocity.X += DECELERATION * dt
			if p.Velocity.X > 0 {
				p.Velocity.X = 0
			}
		}
	}

	// Limita velocidade máxima
	if p.Velocity.X > MAX_SPEED {
		p.Velocity.X = MAX_SPEED
	} else if p.Velocity.X < -MAX_SPEED {
		p.Velocity.X = -MAX_SPEED
	}
}

// HandleJump gerencia a mecânica de pulo
func (p *Player) HandleJump(dt float32) {
	jumpPressed := rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp)
	jumpHeld := rl.IsKeyDown(rl.KeySpace) || rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)

	if jumpPressed {
		p.JumpBufferCounter = JUMP_BUFFER_TIME
	}

	canJump := p.OnGround || p.CoyoteTimeCounter > 0

	if p.JumpBufferCounter > 0 && canJump {
		p.Velocity.Y = p.JumpForce * dt
		p.IsJumping = true
		p.JumpBufferCounter = 0
		p.CoyoteTimeCounter = 0
		p.OnGround = false
		p.CanDoubleJump = p.HasDoubleJump
	}

	// Duplo pulo
	if jumpPressed && !p.OnGround && p.CanDoubleJump && p.HasDoubleJump {
		p.Velocity.Y = -DOUBLE_JUMP_FORCE * dt
		p.IsDoubleJumping = true
		p.CanDoubleJump = false
	}

	// Pulo variável (soltar o botão encurta o pulo)
	if p.Velocity.Y < 0 && !jumpHeld {
		p.Velocity.Y *= 0.5
	}
}

// TryDash tenta iniciar um dash
func (p *Player) TryDash() {
	if rl.IsKeyPressed(rl.KeyLeftShift) && p.HasDash && p.DashCooldown <= 0 {
		p.IsDashing = true
		p.DashTimer = DASH_DURATION
		p.DashDirection = p.FacingDirection

		if p.DashDirection == 0 {
			p.DashDirection = 1
		}
	}
}

// UpdateDash atualiza o estado do dash
func (p *Player) UpdateDash(dt float32) {
	p.Velocity.X = float32(p.DashDirection) * DASH_SPEED * dt
	p.Velocity.Y = 0
	p.DashTimer--

	if p.DashTimer <= 0 {
		p.IsDashing = false
		p.DashCooldown = DASH_COOLDOWN
	}
}

// CheckCollisions detecta e resolve colisões
func (p *Player) CheckCollisions(platforms []objects.Platform) {
	p.OnGround = false
	playerRect := rl.NewRectangle(p.Position.X, p.Position.Y, p.Size.X, p.Size.Y)

	for _, platform := range platforms {
		platformRect := rl.NewRectangle(platform.Position.X, platform.Position.Y, platform.Size.X, platform.Size.Y)
		if rl.CheckCollisionRecs(playerRect, platformRect) {
			p.ResolveCollision(platformRect)
		}
	}

	p.ApplyScreenBounds()
}

// ResolveCollision resolve a colisão com uma plataforma
func (p *Player) ResolveCollision(platform rl.Rectangle) {
	overlapLeft := (p.Position.X + p.Size.X) - platform.X
	overlapRight := (platform.X + platform.Width) - p.Position.X
	overlapTop := (p.Position.Y + p.Size.Y) - platform.Y
	overlapBottom := (platform.Y + platform.Height) - p.Position.Y

	minOverlap := float32(math.Min(float64(overlapLeft),
		math.Min(float64(overlapRight),
			math.Min(float64(overlapTop), float64(overlapBottom)))))

	switch minOverlap {
	case overlapTop:
		if p.Velocity.Y >= 0 {
			p.Position.Y = platform.Y - p.Size.Y
			p.Velocity.Y = 0
			p.OnGround = true
			p.IsJumping = false
			p.IsDoubleJumping = false
		}
	case overlapBottom:
		if p.Velocity.Y <= 0 {
			p.Position.Y = platform.Y + platform.Height
			p.Velocity.Y = 0
		}
	case overlapLeft:
		if p.Velocity.X >= 0 {
			p.Position.X = platform.X - p.Size.X
			p.Velocity.X = 0
		}
	case overlapRight:
		if p.Velocity.X <= 0 {
			p.Position.X = platform.X + platform.Width
			p.Velocity.X = 0
		}
	}
}

// ApplyScreenBounds mantém o jogador dentro da tela
func (p *Player) ApplyScreenBounds() {
	// Limites horizontais
	if p.Position.X < 0 {
		p.Position.X = 0
		p.Velocity.X = 0
	} else if p.Position.X+p.Size.X > SCREEN_WIDTH {
		p.Position.X = SCREEN_WIDTH - p.Size.X
		p.Velocity.X = 0
	}

	// Limites verticais
	if p.Position.Y < 0 {
		p.Position.Y = 0
		p.Velocity.Y = 0
	} else if p.Position.Y+p.Size.Y > SCREEN_HEIGHT {
		p.Position.Y = SCREEN_HEIGHT - p.Size.Y
		p.Velocity.Y = 0
		p.OnGround = true
		p.IsJumping = false
		p.IsDoubleJumping = false
	}
}

// TakeDamage aplica dano ao jogador
func (p *Player) TakeDamage() {
	if !p.Invincible {
		p.Lives--
		p.Invincible = true
		p.InvincibleTimer = INVINCIBLE_TIME
		p.Velocity.Y = -300
	}
}
