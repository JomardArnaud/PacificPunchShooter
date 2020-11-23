package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/veandco/go-sdl2/sdl"
)

//IOpScene is interface for scene
type IOpScene interface {
	GetFileConfig() string
	Init(OpGameConfig)
	Reset(OpGameConfig)
	Event(OpGameConfig, sdl.Event) string
	Update(OpGameConfig, float64) string
	Draw(screen *ebiten.Image)
	PassInfoToNextScene(IOpScene)
	Clean()
}

//OpSceneManager contains and manage scene's game
type OpSceneManager struct {
	idScene, idPrevScene string
	allScene             map[string]IOpScene //be carefull to not use twice the same name for a scene
}

func (manager *OpSceneManager) pushScene(gameInfo OpGameConfig, nSc IOpScene) {
	manager.allScene[nSc.GetFileConfig()] = nSc
	manager.allScene[nSc.GetFileConfig()].Init(gameInfo)
}

func (manager *OpSceneManager) init(gameInfo OpGameConfig) {
	infoManager := NewOpInfoParser(gameInfo.PathConfig + "managerScene.json")

	manager.idScene = infoManager.Blocks["start"].Info["startingScene"]
	manager.allScene = make(map[string]IOpScene)
}

func (manager *OpSceneManager) event(gameInfo OpGameConfig, event sdl.Event) {
	manager.idPrevScene = manager.idScene
	manager.idScene = manager.allScene[manager.idScene].Event(gameInfo, event)
	if manager.idPrevScene != manager.idScene {
		manager.allScene[manager.idPrevScene].PassInfoToNextScene(manager.allScene[manager.idScene])
		manager.allScene[manager.idScene].Reset(gameInfo)
	}
}

func (manager *OpSceneManager) update(gameInfo OpGameConfig, elapsedTime float64) {
	manager.idPrevScene = manager.idScene
	manager.idScene = manager.allScene[manager.idScene].Update(gameInfo, elapsedTime)
	if manager.idPrevScene != manager.idScene {
		manager.allScene[manager.idPrevScene].PassInfoToNextScene(manager.allScene[manager.idScene])
		manager.allScene[manager.idScene].Reset(gameInfo)
	}
}

func (manager *OpSceneManager) draw(screen *ebiten.Image) {
	manager.allScene[manager.idScene].Draw(screen)
}
