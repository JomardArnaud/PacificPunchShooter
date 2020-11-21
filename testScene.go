package main

import "github.com/hajimehoshi/ebiten"

type testScene struct {
	sceneID string
}

func (sc *testScene) GetFileConfig() string {
	return sc.sceneID
}

func (sc *testScene) Init(OpGameConfig) {
	sc.sceneID = "testScene"
}

func (sc *testScene) Reset(OpGameConfig) {

}

func (sc *testScene) Event(OpGameConfig) string {
	return sc.sceneID
}

func (sc *testScene) Update(OpGameConfig, float64) string {
	return sc.sceneID
}

func (sc *testScene) Draw(screen *ebiten.Image) {

}

func (sc *testScene) PassInfoToNextScene(IOpScene) {

}

func (sc *testScene) Clean() {

}
