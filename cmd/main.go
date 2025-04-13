package main

import "github.com/codinomello/arcade-go/core/app"

func main() {
	game := app.NewGame()

	game.Run()

	game.Shutdown()
}
