package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/veandco/go-sdl2/sdl"
)

//OpGameConfig all info may be usefull for other companent
type OpGameConfig struct {
	PathConfig string
	DeadZone   float64
	Buffer     float64
	PosWindow  OpVector2i
	SizeWindow OpVector2i
	Input      OpInput
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
	info.Input.init(info.DeadZone, info.Buffer)
}

//PushScenes in the OpSceneManager
func (g *Game) PushScenes(scenes ...IOpScene) {
	for _, scene := range scenes {
		g.scManager.pushScene(g.config, scene)
	}
}

//Init all setup for running the game
func (g *Game) init() {
	defer func() {
		g.inited = true
	}()
	err := sdl.Init(sdl.INIT_GAMECONTROLLER)
	if err != nil {
		fmt.Println(err)
	}
	g.config.PathConfig = "./assets/config/"
	gameParser := NewOpInfoParser(g.config.PathConfig + "game.json")
	g.config.Init(gameParser)
	g.scManager.init(g.config)
	g.PushScenes(&testScene{sceneID: "testScene"})
	g.config.SizeWindow.X, g.config.SizeWindow.Y = ebiten.WindowSize()
	g.config.PosWindow.X, g.config.PosWindow.Y = ebiten.WindowPosition()
}

//Update the game
func (g *Game) Update() error {
	if !g.inited {
		g.init()
		g.clock = time.Now()
	}

	elapsedTime := time.Since(g.clock).Seconds()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		g.scManager.event(g.config, event)
	}
	g.scManager.update(g.config, elapsedTime)
	g.clock = time.Now()
	return nil
}

//Draw on main screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.scManager.draw(screen)
}

//CleanUpSDL stuff
func (g *Game) CleanUpSDL() {
	if len(g.config.Input.Gamepads) > 0 {
		for i := sdl.NumJoysticks() - 1; i >= 0; i-- {
			g.config.Input.Gamepads[i].Close()
		}
	}
	sdl.Quit()
}

//Layout don't understand what is doing
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if !g.inited {
		return 1, 1
	}
	return g.config.SizeWindow.X, g.config.SizeWindow.Y
}
