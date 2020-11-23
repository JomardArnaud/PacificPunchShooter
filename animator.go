package main

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//OpAnimator manage a Texture to draw a suite's sprite on renderer
type OpAnimator struct {
	OpInfo2d
	AnimeImg   *ebiten.Image
	animations map[string]*OpAnimation
	State      string
}

//OpAnimation struct containe animation
type OpAnimation struct {
	imgRect      image.Rectangle
	sizeCase     OpVector2i
	state        string
	nbAnim       int
	idFrame      int
	line         int
	startingLine int
	framePerLine int
	timePerFrame float64
	timeFrame    float64
}

//InitFromFile push a new animation into the animator
func (anim *OpAnimator) InitFromFile(blockAnimator OpBlock) {
	tmpImage, _, err := ebitenutil.NewImageFromFile(blockAnimator.Info["path"])
	Check(err)
	anim.AnimeImg = ebiten.NewImageFromImage(tmpImage)
	anim.OpInfo2d.InitFromFile(blockAnimator)
	anim.State = blockAnimator.Info["state"]
	anim.animations = make(map[string]*OpAnimation)
	for key, block := range blockAnimator.Blocks["animations"].Blocks {
		anim.animations[key] = NewAnimationFromFile(block)
	}
}

//Update the animation's state
func (anim *OpAnimator) Update(elapsedTime float64) {
	anim.animations[anim.State].Update(elapsedTime)
}

//Draw the image on screen
func (anim *OpAnimator) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(anim.Size.X/float64(anim.animations[anim.State].sizeCase.X), anim.Size.Y/float64(anim.animations[anim.State].sizeCase.Y))
	op.GeoM.Translate(anim.Pos.X, anim.Pos.Y)
	screen.DrawImage(anim.AnimeImg.SubImage(anim.animations[anim.State].imgRect).(*ebiten.Image), op)
}

//Update the rectangle animation
func (a *OpAnimation) Update(elapsedTime float64) {
	a.timeFrame = Clamp(a.timeFrame-elapsedTime, 0.0, a.timePerFrame)

	if a.timeFrame == 0.0 {
		a.idFrame++
		if a.idFrame%a.framePerLine == 0 {
			a.line++
		}
		if a.idFrame == a.nbAnim {
			a.idFrame = 0
			a.line = a.startingLine
		}
		a.timeFrame = a.timePerFrame
	}
	a.imgRect.Min = image.Point{
		X: (a.idFrame % a.framePerLine) * a.sizeCase.X,
		Y: a.line * a.sizeCase.Y}
	a.imgRect.Max = image.Point{
		X: a.imgRect.Min.X + a.sizeCase.X,
		Y: a.imgRect.Min.Y + a.sizeCase.Y}

}

//NewAnimationFromFile create
func NewAnimationFromFile(blockAnimation OpBlock) *OpAnimation {
	return &OpAnimation{
		sizeCase:     OpSetOpVector2i(blockAnimation.Info["sizeCase"]),
		nbAnim:       OpSetInt(blockAnimation.Info["nbAnim"]),
		startingLine: OpSetInt(blockAnimation.Info["startingLine"]),
		line:         OpSetInt(blockAnimation.Info["startingLine"]),
		framePerLine: OpSetInt(blockAnimation.Info["framePerline"]),
		timePerFrame: OpSetFloat(blockAnimation.Info["timePerFrame"]),
		timeFrame:    OpSetFloat(blockAnimation.Info["timePerFrame"]),
	}
}

//InitFromFile push a new animation into the animator
func (a *OpAnimation) InitFromFile(blockAnimation OpBlock) {
	a.sizeCase = OpSetOpVector2i(blockAnimation.Info["sizeCase"])
	a.nbAnim = OpSetInt(blockAnimation.Info["nbAnim"])
	a.startingLine = OpSetInt(blockAnimation.Info["startingLine"])
	a.line = a.startingLine
	a.framePerLine = OpSetInt(blockAnimation.Info["framePerline"])
	a.timePerFrame = OpSetFloat(blockAnimation.Info["timePerFrame"])
	a.timeFrame = a.timePerFrame
}

//AddAnimations push a new animation into the animator
// func (anim *OpAnimator) AddAnimations()
