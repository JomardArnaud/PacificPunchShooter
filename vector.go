package main

import "math"

//OpVector2f is a struct contains 2 float and bunch of method
type OpVector2f struct {
	X, Y float64
}

//OpVector2i is a struct contains 2 int and bunch of method
type OpVector2i struct {
	X, Y int
}

//AddForce to the pos with another OpVector2f
func (v *OpVector2f) AddForce(v2 OpVector2f) {
	v.X += v2.X
	v.Y += v2.Y
}

//AddForces to the pos with two float64
func (v *OpVector2f) AddForces(x, y float64) {
	v.X += x
	v.Y += y
}

//MulForce apply a multiplacator to the fields
func (v *OpVector2f) MulForce(a float64) {
	v.X *= a
	v.Y *= a
}

//NormalizeVector normalize the force of vector
func (v *OpVector2f) NormalizeVector() {
	m := Magnitude(*v)
	if m != 0 {
		v.MulForce(1 / m) //dunno if it not better to v.X /= m && v.Y /= m
	}
}

//RotateVector rotate the vector with the angle in param
func (v *OpVector2f) RotateVector(angle float64) {
	v.X = math.Cos(angle)*v.X - math.Sin(angle)*v.Y
	v.Y = math.Sin(angle)*v.X + math.Cos(angle)*v.Y
}

//AddVector return the sum of two OpVector2f
func AddVector(v, v2 OpVector2f) OpVector2f {
	return (OpVector2f{v.X + v2.X, v.Y + v2.Y})
}

//MulVector return the mul of two OpVector2f
func MulVector(v OpVector2f, a float64) OpVector2f {
	return (OpVector2f{v.X * a, v.Y * a})
}
