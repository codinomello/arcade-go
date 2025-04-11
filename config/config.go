package config

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	ScreenWidth  int32
	ScreenHeight int32
}

func LoadWindowConfig() *Config {
	// Load the base configuration for the game
	config := &Config{
		ScreenWidth:  800,
		ScreenHeight: 600,
	}

	return config
}

func LoadWindowIcon(icon, format string) {
	iconPath := fmt.Sprintf("../images/icons/%s.%s", icon, format)
	loadedIcon := rl.LoadImage(iconPath)
	defer rl.UnloadImage(loadedIcon)
	rl.SetWindowIcon(*loadedIcon)
}
