package main

import "math"

//OpInfo2d is a struct to manage collider and Step of circle in 2D
type OpInfo2d struct {
	Pos     OpVector2f
	Size    OpVector2f
	Force   OpVector2f
	Inertia float64
	Step    float64
}

//InitFromFile config
func (info *OpInfo2d) InitFromFile(infoBlock OpBlock) {
	info.Pos = OpSetOpVector2f(infoBlock.Info["pos"])
	info.Size = OpSetOpVector2f(infoBlock.Info["size"])
	info.Force = OpSetOpVector2f(infoBlock.Info["force"])
	info.Inertia = OpSetFloat(infoBlock.Info["inertia"])
	info.Step = OpSetFloat(infoBlock.Info["step"])
}

//ApplyForce on the Object using the elapsedTime
func (info *OpInfo2d) ApplyForce(elapsedTime float64) {
	info.Pos.AddForces(info.Force.X*elapsedTime, info.Force.Y*elapsedTime)
	info.Force.MulForce(info.Inertia)
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
