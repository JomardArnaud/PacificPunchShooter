package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//OpSprite contains info to isplay image
type OpSprite struct {
	Image  *ebiten.Image
	Filter ebiten.Filter
	OpInfo2d
}

//Update sprite
func (s *OpSprite) Update(elapsedTime float64) {
	s.OpInfo2d.ApplyForce(elapsedTime)
}

//Init all info's OpSprite from info of config file
func (s *OpSprite) Init(infoSpriteBlock OpBlock) {
	tmpImage, _, err := ebitenutil.NewImageFromFile(infoSpriteBlock.Info["path"])
	Check(err)
	s.Image = ebiten.NewImageFromImage(tmpImage)
	s.Filter = ebiten.Filter(OpSetInt(infoSpriteBlock.Info["filter"]))
	s.OpInfo2d = OpInfo2d{
		Pos:     OpSetOpVector2f(infoSpriteBlock.Info["pos"]),
		Size:    OpSetOpVector2f(infoSpriteBlock.Info["size"]),
		Force:   OpSetOpVector2f(infoSpriteBlock.Info["force"]),
		Inertia: OpSetFloat(infoSpriteBlock.Info["inertia"]),
		Step:    OpSetFloat(infoSpriteBlock.Info["step"]),
	}
}

// Draw draws the sprite.
func (s *OpSprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Image.Size()
	op.GeoM.Scale(s.Size.X/float64(w), s.Size.Y/float64(h))
	op.GeoM.Translate(s.Pos.X, s.Pos.Y)
	screen.DrawImage(s.Image, op)
}
