package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type testScene struct {
	sceneID string
	//sprite  OpSprite
	test OpAnimator
}

func (sc *testScene) GetFileConfig() string {
	return sc.sceneID
}

func (sc *testScene) Init(infoGame OpGameConfig) {
	tmpInfoScene := NewOpInfoParser(infoGame.PathConfig + "testScene.json")
	//sc.sprite.Init(tmpInfoScene.Blocks["sprites"].Blocks["cursor"])
	sc.test.InitFromFile(tmpInfoScene.Blocks["animators"].Blocks["jimmyAnimator"])
	sc.sceneID = "testScene"
}

func (sc *testScene) Reset(infoGame OpGameConfig) {

}

func (sc *testScene) Event(infoGame OpGameConfig, event sdl.Event) string {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		fmt.Println("quit ")
	case *sdl.ControllerDeviceEvent:
		switch t.GetType() {
		case sdl.CONTROLLERDEVICEADDED:
			infoGame.Input.pushGamepad()
		case sdl.CONTROLLERDEVICEREMOVED:
			if len(infoGame.Input.Gamepads) == 1 {
				infoGame.Input.Gamepads = nil
			} else {
				infoGame.Input.deleteGamepad(t.Which)
			}
		}
	// case *sdl.ControllerAxisEvent:
	// 	//standby i need to take time to think about how handle axis movement
	// 	game.Input.Gamepads[t.Which].Axis[t.Axis] = int(t.Value)
	// 	break
	case *sdl.ControllerButtonEvent:
		if t.State != 0 { //push a new buffer only if the button is push
		}
	case *sdl.KeyboardEvent:
		if t.GetType() == sdl.KEYUP {
		}
		if t.GetType() == sdl.KEYDOWN { // value t.Keysym.Sym
			if t.Repeat == 0 {
			}
		}
	}
	return sc.sceneID
}

func (sc *testScene) Update(infoGame OpGameConfig, elapsedTime float64) string {
	//sc.sprite.Update(elapsedTime)
	sc.test.Update(elapsedTime)
	return sc.sceneID
}

func (sc *testScene) Draw(screen *ebiten.Image) {
	//sc.sprite.Draw(screen)
	sc.test.Draw(screen)
}

func (sc *testScene) PassInfoToNextScene(IOpScene) {

}

func (sc *testScene) Clean() {

}
