package main

import (
	"fmt"
	"log"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
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

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(screenWidth, screenHeight, "get real")

	level, err := tiled.LoadFromFile("levels/level1.tmx")
	if err != nil {
		log.Fatalf("failed loading tile for rendering: %s", err)
	}

	renderer, err := render.NewRenderer(level)
	if err != nil {
		log.Fatalf("map unsupported for rendering: %s", err)
	}

	err = renderer.RenderVisibleLayers()
	if err != nil {
		log.Fatalf("layer unsupported for rendering: %s", err)
	}

	img := rl.NewImageFromImage(renderer.Result)
	defer rl.UnloadImage(img)

	bg := rl.LoadTextureFromImage(img)
	defer rl.UnloadTexture(bg)

	// Load image and rotate
	carImg := rl.LoadImage("assets/Cars/car_black_3.png")
	defer rl.UnloadImage(carImg)
	rl.ImageRotateCW(carImg)

	carTexture := rl.LoadTextureFromImage(carImg)

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

		x := int(car.X / (128 * 2)) // 2 is the map scale, 128 is the tile size
		y := int(car.Y / (128 * 2)) // 2 is the map scale, 128 is the tile size

		camera.Target = car.Vector2

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawTextureEx(bg, rl.Vector2{X: 0, Y: 0}, 0, 2, rl.White)
		car.Draw()

		currentTile := level.Layers[1].Tiles[y*level.Width+x]
		for _, objGroup := range level.Tilesets[0].Tiles[currentTile.ID].ObjectGroups {
			for _, obj := range objGroup.Objects {
				// We don't care about this object if it's a property
				if len(obj.Properties) != 0 {
					continue
				}
				start := rl.Vector2{X: (float32(obj.X) + float32(x)*128) * 2, Y: (float32(obj.Y) + float32(y)*128) * 2}
				rl.DrawCircleV(start, 10, rl.Pink)

				if len(obj.Ellipses) == 1 {
					// TODO: render Ellipse
				}

				if len(obj.PolyLines) == 1 {
					// TODO: render PolyLine
				}

				if len(obj.Polygons) == 1 {
					// for _, point := range *obj.Polygons[0].Points {
					// 	end := rl.Vector2{X: float32(point.X) + float32(x)*128*2, Y: float32(point.Y) + float32(y)*128*2}
					// 	rl.DrawLineV(start, end, rl.Gray)
					// 	start = end
					// }

					// rl.DrawLineV(origin, rl.Vector2{float32(points[0].X) + float32(x)*128*2, float32(points[0].Y) + float32(y)*128*2}, rl.Gray)
					// rl.Draw
				}
			}
		}

		rl.EndMode2D()

		rl.DrawText("Press R to restart", 5, 5, 25, rl.Black)
		rl.DrawText("Use WASD to move", 5, 35, 25, rl.Black)
		rl.DrawText(fmt.Sprintf("%d", int(car.Speed)), 1024-80, 768-50, 35, rl.Black)
		rl.DrawText(fmt.Sprintf("%d,%d", x, y), 5, 768-50, 35, rl.Black)
		// rl.DrawText(fmt.Sprintf("currentTile.ID = %d", currentTile.ID), 5, 768-50, 35, rl.Black)

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
