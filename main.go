package main

import (
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

const squareSize = 20
const limit = 500

func (c Car) Draw() {
	rl.DrawCircleV(c.FrontWheel, 10, rl.Green)
	rl.DrawLineEx(c.BackWheel, c.FrontWheel, 50, rl.Pink)
}

func (c *Car) Update(dt float64) {
	frontWheel := rl.Vector2Add(
		c.Vector2,
		rl.Vector2Scale(
			rl.Vector2{
				X: float32(math.Cos(deg2Rad(c.Heading))),
				Y: float32(math.Sin(deg2Rad(c.Heading))),
			},
			float32(c.WheelBase/2),
		),
	)

	backWheel := rl.Vector2Subtract(
		c.Vector2,
		rl.Vector2Scale(
			rl.Vector2{
				X: float32(math.Cos(deg2Rad(c.Heading))),
				Y: float32(math.Sin(deg2Rad(c.Heading))),
			},
			float32(c.WheelBase/2),
		),
	)

	backWheel = rl.Vector2Add(
		backWheel,
		rl.Vector2Scale(
			rl.Vector2{
				X: float32(math.Cos(deg2Rad(c.Heading))),
				Y: float32(math.Sin(deg2Rad(c.Heading))),
			},
			float32(c.Speed*dt),
		),
	)

	frontWheel = rl.Vector2Add(
		frontWheel,
		rl.Vector2Scale(
			rl.Vector2{
				X: float32(math.Cos(deg2Rad(c.Heading + c.SteeringAngle))),
				Y: float32(math.Sin(deg2Rad(c.Heading + c.SteeringAngle))),
			},
			float32(c.Speed*dt),
		),
	)

	c.Heading = rad2Deg(math.Atan2(deg2Rad(float64(frontWheel.Y-backWheel.Y)), deg2Rad(float64(frontWheel.X-backWheel.X))))
	c.BackWheel = backWheel
	c.FrontWheel = frontWheel
	c.Vector2 = rl.Vector2Scale(
		rl.Vector2Add(
			frontWheel,
			backWheel,
		), 0.5)
}

var (
	maxBuildings int = 500
)

func main() {
	screenWidth := int32(1024)
	screenHeight := int32(768)

	rl.InitWindow(screenWidth, screenHeight, "get real")

	car := Car{WheelBase: 70, Vector2: rl.Vector2{X: 0, Y: 0}, Heading: 0, SteeringAngle: 0}

	buildings := make([]rl.Rectangle, maxBuildings)
	buildColors := make([]rl.Color, maxBuildings)

	spacing := float32(0)

	for i := 0; i < maxBuildings; i++ {
		r := rl.Rectangle{}
		r.Width = float32(rl.GetRandomValue(50, 200))
		r.Height = float32(rl.GetRandomValue(100, 800))
		r.Y = float32(screenHeight) - 130 - r.Height
		r.X = -6000 + spacing

		spacing += r.Width

		c := rl.NewColor(byte(rl.GetRandomValue(200, 240)), byte(rl.GetRandomValue(200, 240)), byte(rl.GetRandomValue(200, 250)), byte(255))

		buildings[i] = r
		buildColors[i] = c
	}

	camera := rl.Camera2D{
		Target:   rl.NewVector2(float32(car.X+20), float32(car.Y+20)),
		Offset:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		Rotation: 0,
		Zoom:     1,
	}

	rl.SetTargetFPS(60)

	prev := time.Now()
	for !rl.WindowShouldClose() {
		// turn := 0.0
		if rl.IsKeyDown(rl.KeyR) {
			car = Car{WheelBase: 70, Vector2: rl.Vector2{X: 0, Y: 0}, Heading: 0, SteeringAngle: 0}
		}

		if rl.IsKeyDown(rl.KeyW) {
			car.Speed += 5
		} else {
			car.Speed *= 0.99
		}
		if rl.IsKeyDown(rl.KeyS) {
			car.Speed -= 5
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
			// turn += 1
			car.SteeringAngle = 30
		}
		if rl.IsKeyDown(rl.KeyA) {
			car.SteeringAngle = -30
		}

		dt := time.Since(prev)
		prev = time.Now()
		car.Update(dt.Seconds())

		camera.Target = car.Vector2

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		// Draw grid lines

		rl.BeginMode2D(camera)
		for i := range buildings {
			top := buildings[i]
			bottom := rl.Rectangle{
				X:      top.X,
				Y:      top.Y + top.Height,
				Height: top.Height,
				Width:  top.Width,
			}
			rl.DrawRectangleRec(top, buildColors[i])
			rl.DrawRectangleRec(bottom, buildColors[i])
		}
		car.Draw()

		rl.EndMode2D()

		rl.DrawText("Press R to restart", 5, 5, 25, rl.Black)
		rl.DrawText("Use WASD to move", 5, 35, 25, rl.Black)

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
