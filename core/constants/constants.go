package constants

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600

	// Físicas básicas de movimento
	BASE_SPEED     = 300.0  // Velocidade base de movimento
	MAX_SPEED      = 400.0  // Velocidade máxima horizontal
	ACCELERATION   = 2000.0 // Taxa de aceleração horizontal
	DECELERATION   = 1800.0 // Taxa de desaceleração (atrito)
	AIR_RESISTANCE = 0.85   // Multiplicador de controle aéreo

	// Constantes de pulo
	GRAVITY           = 1200.0 // Força da gravidade
	JUMP_FORCE        = 550.0  // Força do pulo inicial
	DOUBLE_JUMP_FORCE = 450.0  // Força do pulo duplo
	JUMP_BUFFER_TIME  = 8      // Frames de buffer para o pulo (permite pressionar antes de tocar o chão)
	COYOTE_TIME       = 6      // Frames extras para pular após sair de uma plataforma

	// Dash
	DASH_SPEED    = 850.0 // Velocidade do dash
	DASH_DURATION = 12    // Duração do dash em frames
	DASH_COOLDOWN = 25    // Tempo de recarga do dash em frames

	// Outros
	INVINCIBLE_TIME = 90 // Tempo de invencibilidade após dano em frames
)
