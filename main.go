package main

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	squareSize = 20
	limit      = 500
	fric       = 0.9
	drag       = 0.02
	mass       = 10
)

func main() {
	screenWidth := int32(1024)
	screenHeight := int32(768)

	rl.InitWindow(screenWidth, screenHeight, "get real")
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	bg := rl.LoadTexture("assets/levels/level1.png")
	defer rl.UnloadTexture(bg)

	carTexture := rl.LoadTexture("assets/cars/car.png")
	defer rl.UnloadTexture(carTexture)

	car := NewCar(float64(carTexture.Width-45.0), carTexture)

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
			car = NewCar(float64(carTexture.Width-45.0), carTexture)
		}

		if rl.IsKeyDown(rl.KeyW) {
			car.Speed += 20
		}
		if rl.IsKeyDown(rl.KeyS) {
			car.Speed -= 20
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

		rl.DrawTextureEx(bg, rl.Vector2{-1024, -768}, 0, 2, rl.White)
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
