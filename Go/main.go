package main

import (
	"math/rand/v2"

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
	Pos    rl.Vector2
	Size   rl.Vector2
	Health float32
	Power  int
	Weapon int
}
type Ground struct {
	Pos     rl.Vector2
	Texture rl.Texture2D
}
type CactusTextures struct {
	Texture rl.Texture2D
}
type WorldCactus struct {
	Pos     rl.Vector2
	Texture rl.Texture2D
}

var RoadSpeed int = 70
var Timer float32 = 0.0
var Offset float32 = 0
var BG1Offset float32 = 0
var dT float32 = 0.0

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "Road Of Fury 2 Clone - Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	Screen := Screen{X: 1920, Y: 1080}

	World := World{X: 11290, Y: 6000}

	Car1 := Car{
		Pos:    rl.NewVector2(World.X, World.Y),
		Size:   rl.NewVector2(400, 300),
		Health: 100,
		Power:  1,
		Weapon: 1,
	}

	Camera := rl.Camera2D{Offset: rl.NewVector2(Screen.X/2, Screen.Y/2), Target: rl.NewVector2(World.X/2, World.Y/2), Rotation: 0, Zoom: 0.17}

	Car_Hud := rl.LoadTexture("./Textures/Car_Hud_Texture.png")

	RoadTexture := rl.LoadTexture("./Textures/Road.png")
	BackGround1 := rl.LoadTexture("./Textures/Background1.png")

	CactusTextures := []CactusTextures{
		{Texture: rl.LoadTexture("./Textures/Cactus/Cactus1.png")},
		{Texture: rl.LoadTexture("./Textures/Cactus/Cactus1.png")},
	}
	Cactus := []WorldCactus{}

	Roads := []Ground{
		{Pos: rl.NewVector2(0, World.Y-1900), Texture: RoadTexture},
		{Pos: rl.NewVector2(1920*5, World.Y-1900), Texture: RoadTexture},
		{Pos: rl.NewVector2(1920*10, World.Y-1900), Texture: RoadTexture},
		{Pos: rl.NewVector2(1920*15, World.Y-1900), Texture: RoadTexture},
	}

	Background1s := []Ground{
		{Pos: rl.NewVector2(0, World.Y-2300), Texture: BackGround1},
		{Pos: rl.NewVector2(1920*5, World.Y-2300), Texture: BackGround1},
		{Pos: rl.NewVector2(1920*10, World.Y-2300), Texture: BackGround1},
		{Pos: rl.NewVector2(1920*15, World.Y-2300), Texture: BackGround1},
	}

	for !rl.WindowShouldClose() {

		dT = rl.GetFrameTime()
		Timer += dT

		if Timer > 10 {
			RoadSpeed += 10
			Timer = 0.0
		}

		Offset -= float32(RoadSpeed)
		BG1Offset -= float32(RoadSpeed) / 1.3

		if Offset <= -(1920 * 4 * 5) {
			Offset = 0
		}
		if BG1Offset <= -(1920 * 4 * 5) {
			BG1Offset = 0
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		rl.BeginMode2D(Camera)

		rl.DrawRectangleV(rl.NewVector2(0, 0), rl.Vector2(World), rl.SkyBlue)
		rl.DrawRectangleV(Car1.Pos, Car1.Size, rl.Red)

		DrawGround(Background1s, BG1Offset) // Render Background 1

		DrawGround(Roads, Offset) // Render Road

		rl.EndMode2D()

		rl.DrawTextureEx(Car_Hud, rl.NewVector2(0, 0), 0.0, 1, rl.White)

		rl.EndDrawing()
	}

}

func DrawGround(Ground []Ground, Offset float32) {
	TotalLength := 1920 * 5 * 4

	for i, G := range Ground {
		G.Pos.X = float32(i*1920*5) + Offset
		rl.DrawTextureEx(G.Texture, rl.NewVector2(G.Pos.X, G.Pos.Y), 0.0, 5, rl.White)
	}

	for i, G := range Ground {
		G.Pos.X = float32(i*1920*5) + Offset + float32(TotalLength)
		rl.DrawTextureEx(G.Texture, rl.NewVector2(G.Pos.X, G.Pos.Y), 0.0, 5, rl.White)
	}

}

func DrawCactus(Cactus []WorldCactus, TextureList []CactusTextures, Offset float32, World World) {
	Roll := rand.IntN(100)

	if len(Cactus)-1 < 4 {
		if Roll > 75 {
			Cactus = append(Cactus, WorldCactus{Pos: rl.NewVector2(World.X+rand.Float32()*100, World.Y-(float32(1930)+rand.Float32()*370)), Texture: TextureList[rand.IntN(1)].Texture})
		}
	}
	if len(Cactus) != 0 {
		
	}
}
