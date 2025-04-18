package config

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	ScreenWidth      float32
	ScreenHeight     float32
	Title            string
	TargetFPS        int32
	Fullscreen       bool
	Resizable        bool
	Borderless       bool
	Undecorated      bool
	VSync            bool
	Icon             string
	IconFormat       string
	Background       string
	BackgroundFormat string
	Font             string
	FontSize         int32
	FontColor        rl.Color
	FontStyle        string
}

func LoadWindowConfig() *Config {
	// Carrega as configurações da janela
	config := &Config{
		ScreenWidth:      800,
		ScreenHeight:     600,
		Title:            "Gopher, Go!",
		TargetFPS:        60,
		Fullscreen:       false,
		Resizable:        true,
		Borderless:       false,
		Undecorated:      false,
		VSync:            true,
		Icon:             "gopher",
		IconFormat:       "png",
		Background:       "background",
		BackgroundFormat: "png",
		Font:             "font",
		FontSize:         20,
		FontColor:        rl.Black,
		FontStyle:        "normal",
	}

	return config
}

func LoadWindowIcon(icon, format string) {
	iconPath := fmt.Sprintf("../images/icons/%s.%s", icon, format)
	loadedIcon := rl.LoadImage(iconPath)
	defer rl.UnloadImage(loadedIcon)
	rl.SetWindowIcon(*loadedIcon)
}
