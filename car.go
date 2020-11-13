package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Car struct {
	rl.Vector2
	// Where you want the car to go
	SteeringAngle float64
	// Where the car is pointing
	Heading    float64
	Speed      float64
	WheelBase  float64
	FrontWheel rl.Vector2
	BackWheel  rl.Vector2
	Texture    rl.Texture2D
}

func NewCar(wb float64, texture rl.Texture2D) *Car {
	return &Car{
		WheelBase: wb,
		Texture:   texture,
		Vector2:   rl.Vector2{X: 0, Y: 0},
	}
}

func (c Car) Draw() {
	w := float32(c.Texture.Width)
	h := float32(c.Texture.Height)

	sin, cos := math.Sincos(deg2Rad(c.Heading))
	offset := rl.Vector2{
		X: float32(cos)*w/2 - float32(sin)*h/2,
		Y: float32(sin)*w/2 + float32(cos)*h/2,
	}

	rl.DrawTextureEx(c.Texture, rl.Vector2Subtract(c.Vector2, offset), float32(c.Heading), 1, rl.White)
}

func (c *Car) Update(dt float64) {
	headingVector := rl.Vector2{
		X: float32(math.Cos(deg2Rad(c.Heading))),
		Y: float32(math.Sin(deg2Rad(c.Heading))),
	}

	c.FrontWheel = rl.Vector2Add(c.Vector2,
		rl.Vector2Scale(
			headingVector,
			float32(c.WheelBase/2),
		),
	)

	c.BackWheel = rl.Vector2Subtract(c.Vector2,
		rl.Vector2Scale(
			headingVector,
			float32(c.WheelBase/2),
		),
	)

	c.BackWheel = rl.Vector2Add(c.BackWheel,
		rl.Vector2Scale(
			headingVector,
			float32(c.Speed*dt),
		),
	)

	c.FrontWheel = rl.Vector2Add(c.FrontWheel,
		rl.Vector2Scale(
			rl.Vector2{
				X: float32(math.Cos(deg2Rad(c.Heading + c.SteeringAngle))),
				Y: float32(math.Sin(deg2Rad(c.Heading + c.SteeringAngle))),
			},
			float32(c.Speed*dt),
		),
	)

	c.Heading = float64(rl.Vector2Angle(c.BackWheel, c.FrontWheel))
	c.Vector2 = rl.Vector2Scale(rl.Vector2Add(c.FrontWheel, c.BackWheel), 0.5)

	// f = m * a
	deacc := fric*10 + (drag * math.Pow(c.Speed, 2) / mass)

	if c.Speed > 0 {
		c.Speed -= deacc * dt
	} else {
		c.Speed += deacc * dt
	}

	// Cap speed
	if c.Speed > limit {
		c.Speed = limit
	}
	if c.Speed < -limit {
		c.Speed = -limit
	}
}
