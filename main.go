package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(1, 1)
	ebiten.SetWindowTitle("First Test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
