package main

import (
	"fmt"
	"math"
	"time"

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
}

const (
	squareSize = 20
	limit      = 500
	fric       = 0.9
	drag       = 0.02
	mass       = 10
)

var (
	carTexture rl.Texture2D
)

func (c Car) Draw() {
	w := float32(carTexture.Width)
	h := float32(carTexture.Height)

	sin, cos := math.Sincos(deg2Rad(c.Heading))
	offset := rl.Vector2{
		X: float32(cos)*w/2 - float32(sin)*h/2,
		Y: float32(sin)*w/2 + float32(cos)*h/2,
	}

	rl.DrawTextureEx(carTexture, rl.Vector2Subtract(c.Vector2, offset), float32(c.Heading), 1, rl.White)
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
}

func main() {
	screenWidth := int32(1024)
	screenHeight := int32(768)

	rl.InitWindow(screenWidth, screenHeight, "get real")
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	bg := rl.LoadTexture("levels/level1.png")
	defer rl.UnloadTexture(bg)

	carTexture = rl.LoadTexture("car.png")
	defer rl.UnloadTexture(carTexture)

	car := Car{WheelBase: float64(carTexture.Width - 10.0), Vector2: rl.Vector2{X: 0, Y: 0}, Heading: 0, SteeringAngle: 0}

	camera := rl.Camera2D{
		Target:   rl.NewVector2(float32(car.X+20), float32(car.Y+20)),
		Offset:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		Rotation: 0,
		Zoom:     1,
	}

	rl.SetTargetFPS(60)

	prev := time.Now()
	for !rl.WindowShouldClose() {
		dt := time.Since(prev)
		prev = time.Now()

		if rl.IsKeyDown(rl.KeyR) {
			car = Car{WheelBase: float64(carTexture.Width - 10.0), Vector2: rl.Vector2{X: 0, Y: 0}, Heading: 0, SteeringAngle: 0}
		}

		// f = m * a
		deacc := fric*10 + (drag * math.Pow(car.Speed, 2) / mass)

		if rl.IsKeyDown(rl.KeyW) {
			car.Speed += 20
		}
		if rl.IsKeyDown(rl.KeyS) {
			car.Speed -= 20
		}

		if car.Speed > 0 {
			car.Speed -= deacc * dt.Seconds()
		} else {
			car.Speed += deacc * dt.Seconds()
		}

		// Cap speed
		if car.Speed > limit {
			car.Speed = limit
		}
		if car.Speed < -limit {
			car.Speed = -limit
		}

		car.SteeringAngle = 0
		if rl.IsKeyDown(rl.KeyD) {
			car.SteeringAngle = 50
		}
		if rl.IsKeyDown(rl.KeyA) {
			car.SteeringAngle = -50
		}

		car.Update(dt.Seconds())

		camera.Target = car.Vector2

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawTexture(bg, 0, 0, rl.White)
		car.Draw()

		rl.EndMode2D()

		rl.DrawText("Press R to restart", 5, 5, 25, rl.Black)
		rl.DrawText("Use WASD to move", 5, 35, 25, rl.Black)
		rl.DrawText(fmt.Sprintf("%d", int(car.Speed)), 1024-80, 768-50, 35, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func deg2Rad(x float64) float64 {
	return x * math.Pi / 180
}

func rad2Deg(x float64) float64 {
	return x * 180 / math.Pi
}
