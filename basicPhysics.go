package main

import "math"

//OpPhysics2d is a struct to manage collider and velocity of circle in 2D
type OpPhysics2d struct {
	Origin   OpVector2f
	Size     OpVector2f
	Force    OpVector2f
	Inertia  float64
	Velocity float64
}

func (p *OpPhysics2d) init(pos, size OpVector2f, inertia, velocity float64) {
	p.Origin = pos
	p.Size = size
	p.Inertia = inertia
	p.Velocity = velocity
}

//ApplyForce on the Object using the elapsedTime
func (p *OpPhysics2d) ApplyForce(elapsedTime float64) {
	p.Origin.AddForces(p.Force.X*elapsedTime, p.Force.Y*elapsedTime)
	p.Force.MulForce(p.Inertia)
}

//simple math stuff

//Clamp a value between the min and max in param
func Clamp(value, min, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

//Magnitude of a vector
func Magnitude(v OpVector2f) float64 {
	return (math.Sqrt((v.X * v.X) + (v.Y * v.Y)))
}

//DotProduct return a float64
func DotProduct(v1 OpVector2f, v2 OpVector2f) float64 {
	return (v1.X * v2.X) + (v1.Y * v2.Y)
}

//AngleVector calculate angle between two point
func AngleVector(v1 OpVector2f, v2 OpVector2f) float64 {
	return math.Atan2(v2.Y-v1.Y, v2.X-v1.X) * (180 / math.Pi)
}

//Distance between two point
func Distance(p0 OpVector2f, p1 OpVector2f) float64 {
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y

	return math.Sqrt(dx*dx + dy*dy)
}

//NormalizeVector return the normalized vector
func NormalizeVector(v OpVector2f) OpVector2f {
	m := Magnitude(v)
	if m == 0 {
		return OpVector2f{}
	}
	return OpVector2f{v.X / m, v.Y / m}
}

//RotateVector return a new Vector after a rotation by the angle in param
func RotateVector(v OpVector2f, angle float64) OpVector2f {
	return OpVector2f{math.Cos(angle)*v.X - math.Sin(angle)*v.Y, math.Sin(angle)*v.X + math.Cos(angle)*v.Y}
}
