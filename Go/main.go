package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Screen struct {
	X float32
	Y float32
}
type World struct {
	X float32
	Y float32
}
type Car struct {
	Pos rl.Vector2
	Size rl.Vector2
	Health float32
	Power int
	Weapon int
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "Road Of Fury 2 Clone - Go")
	defer rl.CloseWindow()

	Screen := Screen{X: 1920,Y: 1080}

	World := World{X: 11290, Y: 6000}

	Car1 := Car{
		Pos: rl.NewVector2(World.X, World.Y), 
		Size: rl.NewVector2(400, 300),
		Health: 100,
		Power: 1,
		Weapon: 1,
	}

	Camera := rl.Camera2D{Offset: rl.NewVector2(Screen.X/2, Screen.Y/2), Target: rl.NewVector2(World.X/2, World.Y/2), Rotation: 0, Zoom: 0.17}

	for !rl.WindowShouldClose() {


		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		rl.BeginMode2D(Camera)

		rl.DrawRectangleV(rl.NewVector2(0,0), rl.Vector2(World), rl.Black)
		rl.DrawRectangleV(Car1.Pos, Car1.Size, rl.Red)

		rl.EndDrawing()
	}
	
}