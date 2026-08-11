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
type Cloud struct {
	Pos rl.Vector2
	Texture rl.Texture2D
	Scale float32
}
type CactusTextures struct {
	Texture rl.Texture2D
}
type WorldCactus struct {
	Pos     rl.Vector2
	Texture rl.Texture2D
}

var RoadSpeed int = 200
var Timer float32 = 0.0
var Offset float32 = 0
var BG1Offset float32 = 0
var BG2Offset float32 = 0
var BG3Offset float32 = 0
var SkyOffset float32 = 0
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
	BackGround2 := rl.LoadTexture("./Textures/Background2.png")
	BackGround3 := rl.LoadTexture("./Textures/Background3.png")
	SkyTexture := rl.LoadTexture("./Textures/Sky.png")  
	CloudTexture := rl.LoadTexture("./Textures/Cloud.png") //542, 163

	Roads := []Ground{
		{Pos: rl.NewVector2(0, World.Y-1900), Texture: RoadTexture},
		{Pos: rl.NewVector2(1920*5, World.Y-1900), Texture: RoadTexture},
		{Pos: rl.NewVector2(1920*10, World.Y-1900), Texture: RoadTexture},
		{Pos: rl.NewVector2(1920*15, World.Y-1900), Texture: RoadTexture},
	}

	Background1s := []Ground{
		{Pos: rl.NewVector2(0, World.Y-2700), Texture: BackGround1},
		{Pos: rl.NewVector2(1920*5, World.Y-2700), Texture: BackGround1},
		{Pos: rl.NewVector2(1920*10, World.Y-2700), Texture: BackGround1},
		{Pos: rl.NewVector2(1920*15, World.Y-2700), Texture: BackGround1},
	}
	Background2s := []Ground{
		{Pos: rl.NewVector2(0, 2300), Texture: BackGround2},
		{Pos: rl.NewVector2(1920*5, 2300), Texture: BackGround2},
		{Pos: rl.NewVector2(1920*10, 2300), Texture: BackGround2},
		{Pos: rl.NewVector2(1920*15, 2300), Texture: BackGround2},
	}
	Background3s := []Ground{
		{Pos: rl.NewVector2(0, 500), Texture: BackGround3},
		{Pos: rl.NewVector2(1920*2, 500), Texture: BackGround3},
		{Pos: rl.NewVector2(1920*10, 500), Texture: BackGround3},
		{Pos: rl.NewVector2(1920*15, 500), Texture: BackGround3},
	}
	Skys := []Ground{
		{Pos: rl.NewVector2(0,-200), Texture: SkyTexture},
		{Pos: rl.NewVector2(1920*5,-200), Texture: SkyTexture},
		{Pos: rl.NewVector2(1920*10,-200), Texture: SkyTexture},
		{Pos: rl.NewVector2(1920*15,-200), Texture: SkyTexture},
	}
	Clouds := []Cloud {
		{Pos: rl.NewVector2(World.X - 3000, 0), Texture: CloudTexture},
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
		BG2Offset -= float32(RoadSpeed) / 2.5
		BG3Offset -= float32(RoadSpeed) / 5.5
		SkyOffset -= float32(RoadSpeed) / 15

		Roll := rand.IntN(100)
		if Roll > 90 {
			Clouds = append(Clouds, Cloud{Pos: rl.NewVector2(World.X, rand.Float32()*200), Texture: CloudTexture, Scale: 1+rand.Float32()*4})
		}

		if Offset <= -(1920 * 4 * 5) {
			Offset = 0
		}
		if BG1Offset <= -(1920 * 4 * 5) {
			BG1Offset = 0
		}
		if BG2Offset <= -(1920 * 4 * 5) {
			BG2Offset = 0
		}
		if BG3Offset <= -(1920*4*5) {
			BG3Offset = 0
		}
		if SkyOffset <= -(1920*4*5) {
			SkyOffset = 0
		}
		
		for i,C := range Clouds {
			if C.Pos.X < -542 {
				Clouds[i] = Clouds[len(Clouds)-1]
				Clouds = Clouds[:len(Clouds)-1]
			}
			Clouds[i].Pos.X = float32(i*542*int(C.Scale)) + SkyOffset+ 10 
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		rl.BeginMode2D(Camera)

		rl.DrawRectangleV(rl.NewVector2(0, 0), rl.Vector2(World), rl.SkyBlue)
		rl.DrawRectangleV(Car1.Pos, Car1.Size, rl.Red)

		DrawGround(Skys, SkyOffset)

		
		DrawGround(Background3s, BG3Offset)
		DrawClouds(Clouds)
		rl.DrawTextureEx(Clouds[0].Texture, Clouds[0].Pos, 0.0, 5, rl.White)

		DrawGround(Background2s, BG2Offset)

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

func DrawClouds(Clouds []Cloud) {
	for _,C := range Clouds {
		rl.DrawTextureEx(C.Texture, C.Pos, 0.0, 1+rand.Float32()*4, rl.White)
	}
}


