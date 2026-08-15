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
	Pos           rl.Vector2
	Texture       rl.Texture2D
	Health        float32
	Power         int
	PowerTime     float32
	PowerCooldown float32
	Weapon        int
	Scale         float32
}
type Ground struct {
	Pos     rl.Vector2
	Texture rl.Texture2D
}
type Cloud struct {
	Pos       rl.Vector2
	Texture   rl.Texture2D
	Scale     float32
	Offset    float32
	RandSpeed float32
}
type CactusTextures struct {
	Texture rl.Texture2D
}
type WorldCactus struct {
	Pos     rl.Vector2
	Texture rl.Texture2D
}

var RoadSpeed float32 = 200
var Timer float32
var Offset float32
var BG1Offset float32
var BG2Offset float32
var BG3Offset float32
var SkyOffset float32
var CloudTimer float32
var Car1OffsetX float32
var Car1OffsetTimer float32
var dT float32

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "Road Of Fury 2 Clone - Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	Screen := Screen{X: 1920, Y: 1080}

	World := World{X: 11290, Y: 6000}

	Camera := rl.Camera2D{Offset: rl.NewVector2(Screen.X/2, Screen.Y/2), Target: rl.NewVector2(World.X/2, World.Y/2), Rotation: 0, Zoom: 0.17}

	Car_Hud := rl.LoadTexture("./Textures/Car_Hud_Texture.png")
	Car1_Texture := rl.LoadTexture("./Textures/car1.png")

	RoadTexture := rl.LoadTexture("./Textures/Road.png")
	BackGround1 := rl.LoadTexture("./Textures/Background1.png")
	BackGround2 := rl.LoadTexture("./Textures/Background2.png")
	BackGround3 := rl.LoadTexture("./Textures/Background3.png")
	SkyTexture := rl.LoadTexture("./Textures/Sky.png")
	CloudTexture := rl.LoadTexture("./Textures/Cloud.png")

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
		{Pos: rl.NewVector2(0, -200), Texture: SkyTexture},
		{Pos: rl.NewVector2(1920*5, -200), Texture: SkyTexture},
		{Pos: rl.NewVector2(1920*10, -200), Texture: SkyTexture},
		{Pos: rl.NewVector2(1920*15, -200), Texture: SkyTexture},
	}
	Clouds := []Cloud{
		{Pos: rl.NewVector2(World.X-3000, 0), Texture: CloudTexture},
	}

	Car1 := Car{Pos: rl.NewVector2(-194, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Nulifier, PowerTime: 10, PowerCooldown: 0.0, Weapon: Minigun, Scale: 5}

	for !rl.WindowShouldClose() {

		dT = rl.GetFrameTime()
		Timer += dT

		if Timer > 10 {
			RoadSpeed += 10
			Timer = 0.0
		}

		Offset -= RoadSpeed
		BG1Offset -= RoadSpeed / 1.3
		BG2Offset -= RoadSpeed / 2.5
		BG3Offset -= RoadSpeed / 5.5
		SkyOffset -= RoadSpeed / 15

		if Offset <= -(1920 * 4 * 5) {
			Offset = 0
		}
		if BG1Offset <= -(1920 * 4 * 5) {
			BG1Offset = 0
		}
		if BG2Offset <= -(1920 * 4 * 5) {
			BG2Offset = 0
		}
		if BG3Offset <= -(1920 * 4 * 5) {
			BG3Offset = 0
		}
		if SkyOffset <= -(1920 * 4 * 5) {
			SkyOffset = 0
		}

		// Car Random Movement Timer -------------------------------------------------------------------------------------------------
		if Car1OffsetTimer <= 0 {
			Roll := rand.IntN(2)
			if Roll == 0 {
				Car1OffsetX = rand.Float32() * float32(rand.IntN(101)+5) * Car1.Scale
			} else {
				Car1OffsetX = -rand.Float32() * float32(rand.IntN(101)+5) * Car1.Scale
			}
			Car1OffsetTimer = 4 + rand.Float32()*12
		} else {
			Car1OffsetTimer -= dT
		}
		//-----------------------------------------------------------------------------------------------------------------------------

		// Car Random Movement Handler ------------------------------------------------------------------------------------------------
		switch {
		case Car1OffsetX > 0:
			Car1.Pos.X += 0.5 * Car1.Scale
			Car1OffsetX -= 0.5 * Car1.Scale
			if Car1OffsetX < 0 {
				Car1OffsetX = 0
			}
		case Car1OffsetX < 0:
			Car1.Pos.X -= 0.5 * Car1.Scale
			Car1OffsetX += 0.5 * Car1.Scale
			if Car1OffsetX > 0 {
				Car1OffsetX = 0
			}
		}
		//------------------------------------------------------------------------------------------------------------------------------

		// Stops car from going offscreen ----------------------------------------------------------------------------------------------
		if Car1OffsetX == 0 && Car1.Pos.X < 0 {
			Car1OffsetX = -Car1.Pos.X
		} else if Car1OffsetX == 0 && Car1.Pos.X > 2000 {
			Car1OffsetX = -Car1.Pos.X
		}
		//------------------------------------------------------------------------------------------------------------------------------

		Clouds = CreateClouds(Clouds, CloudTexture, RoadSpeed, World) // BROKEN STUFF AAAAHHHHH

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		rl.BeginMode2D(Camera)

		rl.DrawRectangleV(rl.NewVector2(0, 0), rl.Vector2(World), rl.SkyBlue)

		DrawGround(Skys, SkyOffset)

		DrawClouds(Clouds) // --------------------------------------------------------------------------- FIX CLOUDS PLZ ---------------------------------------------------------------------------------------------------------
		DrawGround(Background3s, BG3Offset)

		DrawGround(Background2s, BG2Offset)

		DrawGround(Background1s, BG1Offset) // Render Background 1

		DrawGround(Roads, Offset) // Render Road

		rl.DrawTextureEx(Car1.Texture, Car1.Pos, 0.0, Car1.Scale, rl.White)

		rl.EndMode2D()

		rl.DrawTextureEx(Car_Hud, rl.NewVector2(0, 0), 0.0, 1, rl.White)

		rl.EndDrawing()
	}

}

// Creates a Moving Foreground or Background depending on the order of use.
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

// Creates and adds Clouds to a list that it also moves, and also removes them. Retuns a Slice of type Struct Cloud.
func CreateClouds(Clouds []Cloud, Texture rl.Texture2D, RoadSpeed float32, World World) []Cloud {
	Roll := rand.IntN(100)
	if Roll > 95 && CloudTimer <= 0 && len(Clouds) < 4 {
		Scale := 1 + rand.Float32()*4
		Clouds = append(Clouds, Cloud{Pos: rl.NewVector2(World.X+542*Scale, rand.Float32()*200), Texture: Texture, Scale: Scale, Offset: 0, RandSpeed: rand.Float32() * 100})
		CloudTimer = 3
	}

	for i := 0; i < len(Clouds)-1; {
		Clouds[i].Offset -= float32(RoadSpeed) / 12
		// if Clouds[i].Offset <= -float32(i*1920*len(Clouds)*int(Clouds[i].Scale)) {
		// if Clouds[i].Offset <= -float32(i*int(World.X)*len(Clouds))*2*(542* Clouds[i].Scale) {
		if Clouds[i].Offset <= -float32(1920*6*Clouds[i].Scale) {
			Clouds[i].Offset = 0
		}
		if Clouds[i].Pos.X < -542 {
			Clouds[i] = Clouds[len(Clouds)-1]
			Clouds = Clouds[:len(Clouds)-1]
		}
		Clouds[i].Pos.X = float32(9542*int(Clouds[i].Scale)) + Clouds[i].Offset + Clouds[i].RandSpeed
		i++
	}

	if CloudTimer >= 0 {
		CloudTimer -= dT
	}
	return Clouds
}

// It's only job is to render the Clouds from the Clouds Slice lol.
func DrawClouds(Clouds []Cloud) {
	for _, C := range Clouds {
		rl.DrawTextureEx(C.Texture, C.Pos, 0.0, C.Scale, rl.White)
	}
}

func RandCarMovementTimer(CarOffsetTimer, CarOffsetX float32, Car Car) (float32, float32) {
	if CarOffsetTimer <= 0 {
		Roll := rand.IntN(2)
		if Roll == 0 {
			CarOffsetX = rand.Float32() * float32(rand.IntN(101)+5) * Car.Scale
		} else {
			CarOffsetX = -rand.Float32() * float32(rand.IntN(101)+5) * Car.Scale
		}
		Car1OffsetTimer = 4 + rand.Float32()*12
	} else {
		Car1OffsetTimer -= dT
	}

	return CarOffsetTimer, CarOffsetX
}
