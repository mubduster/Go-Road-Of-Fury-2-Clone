package main

import (
	// "fmt"
	"math"
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
	OffsetSpeed float32
}
type Gun struct {
	PosBody rl.Vector2
	PosGun rl.Vector2
	Angle	float32
	TextureBody rl.Texture2D
	TextureGun rl.Texture2D
	PrevGunAngle float32
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
var Car2OffsetTimer float32
var Car2OffsetX float32
var Car3OffsetTimer float32
var Car3OffsetX float32
var Gun1Angle float64
var dT float32

func main() {

	// Initialize Game Window ----------------------------------------------------
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "Road Of Fury 2 Clone - Go")
	defer rl.CloseWindow()
	//----------------------------------------------------------------------------
	
	rl.SetTargetFPS(60) // Locking FPS for now so that i can test the game without having to implement delta time. 

	// Loading Textures -----------------------------------------------------------------------------------
	Car_Hud := rl.LoadTexture("./Textures/Car_Hud_Texture.png")
	Car1_Texture := rl.LoadTexture("./Textures/car1.png")

	RoadTexture := rl.LoadTexture("./Textures/Road.png")
	BackGround1 := rl.LoadTexture("./Textures/Background1.png")
	BackGround2 := rl.LoadTexture("./Textures/Background2.png")
	BackGround3 := rl.LoadTexture("./Textures/Background3.png")
	SkyTexture := rl.LoadTexture("./Textures/Sky.png")
	CloudTexture := rl.LoadTexture("./Textures/Cloud.png")
	MiniGunBody := rl.LoadTexture("./Textures/Guns/MiniGunBase.png")
	MiniGunTexture := rl.LoadTexture("./Textures/Guns/MiniGun.png")
	//-----------------------------------------------------------------------------------------------------

	// Create Structs and entities --------------------------------------------------------------------------------------------------------------------------------------------
	Screen := Screen{X: 1920, Y: 1080}
	
	World := World{X: 11290, Y: 6000}

	Camera := rl.Camera2D{Offset: rl.NewVector2(Screen.X/2, Screen.Y/2), Target: rl.NewVector2(World.X/2, World.Y/2), Rotation: 0, Zoom: 0.17}

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

	Car1 := Car{Pos: rl.NewVector2(-694, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Nulifier, PowerTime: 10, PowerCooldown: 0.0, Weapon: Minigun, Scale: 5, OffsetSpeed: 1}
	Car2 := Car{Pos: rl.NewVector2(-394, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Sat, PowerTime: 10, PowerCooldown: 0.0, Weapon: PulseGun, Scale: 5, OffsetSpeed: 1.5}
	Car3 := Car{Pos: rl.NewVector2(-194, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Emp, PowerTime: 10, PowerCooldown: 0.0, Weapon: Laser, Scale: 5, OffsetSpeed: 5}

	Gun1 := Gun{PosBody: rl.NewVector2(Car1.Pos.X+300, World.Y-3000), PosGun: rl.NewVector2(Car1.Pos.X+300-(16*3),World.Y-2350), Angle: 0, TextureBody: MiniGunBody, TextureGun: MiniGunTexture, PrevGunAngle: 0}
	//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------

	// Game Loop ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	for !rl.WindowShouldClose() {

		dT = rl.GetFrameTime()
		Mouse := rl.GetMousePosition()
		WorldMouse := rl.GetScreenToWorld2D(Mouse, Camera)

		// Position the Guns on the cars ------------------------------
		Gun1.PosBody = Car1.Pos.Add(Car1Minigun.GunBody)
		Gun1.PosGun = Car1.Pos.Add(Car1Minigun.GunGun)
		//-------------------------------------------------------------
		// Gun Following Cursor handler -------------------------------
		Gun1.Angle, Gun1.PrevGunAngle = GunFollow(WorldMouse, Gun1)
		//-------------------------------------------------------------

		// Timer to increase the speed -------------------------
		Timer += dT
		if Timer > 10 {
			RoadSpeed += 10
			Timer = 0.0
		}
		//------------------------------------------------------

		// Map Speeds -------------------------------------------
		Offset -= RoadSpeed
		BG1Offset -= RoadSpeed / 1.3
		BG2Offset -= RoadSpeed / 2.5
		BG3Offset -= RoadSpeed / 5.5
		SkyOffset -= RoadSpeed / 15
		//-------------------------------------------------------

		// Border end check -------------------------------------
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
		//-------------------------------------------------------

		// Car Random Movement Timer -------------------------------------------------------------------------------------------------
		Car1OffsetTimer, Car1OffsetX, Car1 = RandCarMovementTimer(Car1OffsetTimer, Car1OffsetX, Car1)
		Car2OffsetTimer, Car2OffsetX, Car2 = RandCarMovementTimer(Car2OffsetTimer, Car2OffsetX, Car2)
		Car3OffsetTimer, Car3OffsetX, Car3 = RandCarMovementTimer(Car3OffsetTimer, Car3OffsetX, Car3)
		//-----------------------------------------------------------------------------------------------------------------------------

		// Car Random Movement Handler ------------------------------------------------------------------------------------------------
		Car1OffsetX, Car1 = RandMove(Car1OffsetX, Car1)
		Car2OffsetX, Car2 = RandMove(Car2OffsetX, Car2)
		Car3OffsetX, Car3 = RandMove(Car3OffsetX, Car3)
		//------------------------------------------------------------------------------------------------------------------------------

		// Stops car from going offscreen ----------------------------------------------------------------------------------------------
		Car1OffsetX = RandMoveCheckBoundry(0, 1200, Car1OffsetX, Car1)
		Car2OffsetX = RandMoveCheckBoundry(700, 1700, Car2OffsetX, Car2)
		Car3OffsetX = RandMoveCheckBoundry(1600, 2500, Car3OffsetX, Car3)
		//------------------------------------------------------------------------------------------------------------------------------

		Clouds = CreateClouds(Clouds, CloudTexture, RoadSpeed, World) // BROKEN STUFF AAAAHHHHH

		// Begin Rendering ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Camera start ------------------------------------------------------------------------------------------------------------------
		rl.BeginMode2D(Camera)

		// Background ---------------------------------------------------------------------------------------------------------------------
		DrawGround(Skys, SkyOffset)

		DrawClouds(Clouds) // <--------------------------------------------------------------------------- FIX CLOUDS PLZ ---------------------------------------------------------------------------------------------------------
		DrawGround(Background3s, BG3Offset)

		DrawGround(Background2s, BG2Offset)

		DrawGround(Background1s, BG1Offset) // Render Background 1
		//---------------------------------------------------------------------------------------------------------------------------------

		DrawGround(Roads, Offset) // Render Road

		
		rl.DrawTextureEx(Car1.Texture, Car1.Pos, 0.0, Car1.Scale, rl.White) // Render Car 1
		rl.DrawTextureEx(Car2.Texture, Car2.Pos, 0.0, Car2.Scale, rl.White)
		rl.DrawTextureEx(Car3.Texture, Car3.Pos, 0.0, Car3.Scale, rl.White)
		
		rl.DrawTexturePro(MiniGunTexture, rl.NewRectangle(0, 0, 39, 7), rl.NewRectangle(Gun1.PosGun.X+85, Gun1.PosGun.Y, 300, 70), rl.NewVector2(85, 30), Gun1.Angle, rl.White)
		rl.DrawTextureEx(Gun1.TextureBody, Gun1.PosBody, 0.0, Car1Minigun.Scale, rl.White)
		
		rl.EndMode2D()
		// Camera end ---------------------------------------------------------------------------------------------------------------------

		rl.DrawTextureEx(Car_Hud, rl.NewVector2(0, 0), 0.0, 1, rl.White) // Render Game Hud

		// rl.DrawTextEx(rl.GetFontDefault(),fmt.Sprintf("Angle: %0.2f , dX: %0.2f , dY: %0.2f", Gun1Angle, dx, dy), rl.NewVector2(10, 10), 50, 10, rl.White)

		rl.EndDrawing()
		//----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	}
	//--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
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

// Handles the Randomness of the movement by making them only move after a set random time has passed.
func RandCarMovementTimer(CarOffsetTimer, CarOffsetX float32, Car Car) (float32, float32, Car) {
	if CarOffsetTimer <= 0 {
		Roll := rand.IntN(2)
		if Roll == 0 {
			CarOffsetX = rand.Float32() * float32(rand.IntN(101)+5) * Car.Scale
		} else {
			CarOffsetX = -rand.Float32() * float32(rand.IntN(101)+5) * Car.Scale
		}
		CarOffsetTimer = 4 + rand.Float32()*12
		Car.OffsetSpeed = 0.5+rand.Float32()*0.5
	} else {
		CarOffsetTimer -= dT
	}

	return CarOffsetTimer, CarOffsetX, Car
}

// Actually Does the Random Movement.
func RandMove(CarOffsetX float32, Car Car) (float32, Car) {
	switch {
	case CarOffsetX > 0:
		Car.Pos.X += Car.OffsetSpeed * Car.Scale
		CarOffsetX -= Car. OffsetSpeed * Car.Scale
		if CarOffsetX < 0 {
			CarOffsetX = 0
		}
	case CarOffsetX < 0:
		Car.Pos.X -= Car.OffsetSpeed * Car.Scale
		CarOffsetX += Car.OffsetSpeed * Car.Scale
		if CarOffsetX > 0 {
			CarOffsetX = 0
		}
	}
	return CarOffsetX, Car
}

// Makes sure that the Car doesn't cross it's boundries.
func RandMoveCheckBoundry(Min, Max, CarOffsetX float32, Car Car) float32 {
	if CarOffsetX == 0 && Car.Pos.X < Min {
		CarOffsetX = Min-Car.Pos.X
	} else if CarOffsetX == 0 && Car.Pos.X > Max {
		CarOffsetX = Car.Pos.X-Max
	}
	return CarOffsetX
}

// Makes the Gun follow the Cursor.
func GunFollow(Mouse rl.Vector2, Gun Gun) (float32, float32) {
		dx := float64(Mouse.X-Gun.PosGun.X)
		dy := float64(Mouse.Y-Gun.PosGun.Y)

		GunAngle := math.Atan2(dy,dx)*180/math.Pi

		if GunAngle > 45 || GunAngle < -135 {
			return float32(Gun.PrevGunAngle), float32(Gun.PrevGunAngle)
		}else if GunAngle > 42 {
			GunAngle = 42 
		}else if GunAngle < -130 {
			GunAngle = -130
		}
	return float32(GunAngle), float32(GunAngle)
}
