package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	closeSDL := &Game{}
	ebiten.SetWindowResizable(true)
	ebiten.MaximizeWindow()
	ebiten.SetWindowTitle("First Test")
	if err := ebiten.RunGame(closeSDL); err != nil {
		log.Fatal(err)
	}
	closeSDL.CleanUpSDL()
}
