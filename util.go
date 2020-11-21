package main

import (
	"log"
	"strconv"
	"strings"
)

//OpCooldown (in seconds) is a little tool to control time will change when i will got a better knowledge of channel and goroutine
type OpCooldown struct {
	Tick, Duration float64
}

//Update the tick using the time elapsed
func (cd *OpCooldown) Update(elapsedTime float64) {
	cd.Tick = Clamp(cd.Tick-elapsedTime, 0.0, cd.Tick)
}

//Reset tick to the duration value
func (cd *OpCooldown) Reset() {
	cd.Tick = cd.Duration
}

//OpSetInt convert a string into a int
func OpSetInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

//OpSetFloat convert a string into a float64
func OpSetFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

//OpSetByte convert a string into a byte
func OpSetByte(str string) byte {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return byte(value)
}

//OpSetBool convert a string into a bool
func OpSetBool(str string) bool {
	value, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

//OpSetOpVector2i convert a string into a OpVector2f
func OpSetOpVector2i(str string) OpVector2i {
	if str != "" {
		s := strings.Split(str, ",")
		return OpVector2i{OpSetInt(s[0]), OpSetInt(s[1])}
	}
	return OpVector2i{}
}

//OpSetOpVector2f convert a string into a OpVector2f
func OpSetOpVector2f(str string) OpVector2f {
	if str != "" {
		s := strings.Split(str, ",")
		return OpVector2f{OpSetFloat(s[0]), OpSetFloat(s[1])}
	}
	return OpVector2f{}
}
