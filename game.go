package main

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

//OpGameConfig all info may be usefull for other companent
type OpGameConfig struct {
	PathConfig              string
	DeadZone, Buffer, Timer float64
	PosWindow, SizeWindow   OpVector2i
}

//Game is the main struct which contains all info to run the game
type Game struct {
	config      OpGameConfig
	elapsedTime float64
	clock       time.Time
	scManager   OpSceneManager
	inited      bool
}

//Init OpGameConfig with info from parser
func (info *OpGameConfig) Init(parser OpInfoParser) {
	info.DeadZone = OpSetFloat(parser.Blocks["config"].Info["deadZone"])
	info.Buffer = OpSetFloat(parser.Blocks["config"].Info["buffer"])
	info.Timer = OpSetFloat(parser.Blocks["config"].Info["timer"])
	info.PosWindow = OpSetOpVector2i(parser.Blocks["config"].Info["posWindow"])
	info.SizeWindow = OpSetOpVector2i(parser.Blocks["config"].Info["sizeWindow"])
}

//PushScenes in the OpSceneManager
func (g *Game) PushScenes(scenes ...IOpScene) {
	for _, scene := range scenes {
		g.scManager.pushScene(g.config, scene)
	}
}

//Init all setup for running the game
func (g *Game) init() {
	g.config.PathConfig = "./assets/config/"
	gameParser := NewOpInfoParser(g.config.PathConfig + "game.json")
	g.config.Init(gameParser)
	g.scManager.init(g.config)
	g.PushScenes(&testScene{sceneID: "testScene"})

	ebiten.SetWindowSize(g.config.SizeWindow.X, g.config.SizeWindow.Y)
}

//Update the game
func (g *Game) Update(screen *ebiten.Image) error {
	if !g.inited {
		g.init()
		g.clock = time.Now()
	}

	elapsedTime := time.Since(g.clock).Seconds()
	//g.scManager
	g.scManager.update(g.config, elapsedTime)
	g.clock = time.Now()
	return nil
}

//Draw on main screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.scManager.draw(screen)
}

//Layout don't understand what is doing
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if !g.inited {
		return 1, 1
	}
	return g.config.SizeWindow.X, g.config.SizeWindow.Y
}
