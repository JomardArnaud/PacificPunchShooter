package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

/*
enum for Keyboard => https://wiki.libsdl.org/SDL_Keycode
	 for button => https://wiki.libsdl.org/SDL_GameControllerButton
	 for axis => https://wiki.libsdl.org/SDL_GameControllerAxis
	 for mouse => https://wiki.libsdl.org/SDL_MouseButtonEvent#Remarks
*/

//OpInput take input from sdl int the event method and stock into a different array (it serve to be able to set to 0 key after they have been used)
type OpInput struct {
	DeadZone float64
	Gamepads []*sdl.GameController
	KeyState []uint8
	//lastInput []OpCooldown each input from last buffer duration
}

func (input *OpInput) init(nDeadZone, buffer float64) {
	input.DeadZone = nDeadZone
	input.KeyState = sdl.GetKeyboardState()
	//gamepad are add through event loop
}

//GetLeftStick return a the position of leftStick of the idPad already deadZone checked and can be normalized
func (input *OpInput) GetLeftStick(ID sdl.GameControllerAxis, normalized bool) OpVector2f {
	tmpStick := OpVector2f{float64(input.Gamepads[ID].Axis(sdl.CONTROLLER_AXIS_LEFTX)), float64(input.Gamepads[ID].Axis(sdl.CONTROLLER_AXIS_LEFTY))}
	if Distance(tmpStick, OpVector2f{0.0, 0.0}) < input.DeadZone {
		tmpStick = OpVector2f{}
	}
	if normalized {
		tmpStick.NormalizeVector()
	}
	return tmpStick
}

//GetRightStick return a the position of rightStick already deadZone checked and can be normalized
func (input *OpInput) GetRightStick(ID sdl.GameControllerAxis, normalized bool) OpVector2f {
	tmpStick := OpVector2f{float64(input.Gamepads[ID].Axis(sdl.CONTROLLER_AXIS_RIGHTX)), float64(input.Gamepads[ID].Axis(sdl.CONTROLLER_AXIS_RIGHTY))}
	if Distance(tmpStick, OpVector2f{0.0, 0.0}) < input.DeadZone {
		tmpStick = OpVector2f{}
	}
	if normalized {
		tmpStick.NormalizeVector()
	}
	return tmpStick
}

func (input *OpInput) pushGamepad() {
	input.Gamepads = append(input.Gamepads, sdl.GameControllerOpen(sdl.NumJoysticks()-1))
}

func (input *OpInput) deleteGamepad(ID sdl.JoystickID) {
	//to avoid memory leaks
	if len(input.Gamepads) > 0 {
		copy(input.Gamepads[ID:], input.Gamepads[ID+1:])
		input.Gamepads[len(input.Gamepads)-1] = nil // or the zero value of T
		input.Gamepads = input.Gamepads[:len(input.Gamepads)-1]
	} else {
		input.Gamepads = nil
	}
}
