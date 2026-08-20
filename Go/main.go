package main

import (
	"fmt"
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
	WheelTexture rl.Texture2D
	Frames int
	Car int
}
type Gun struct {
	PosBody rl.Vector2
	PosGun rl.Vector2
	Angle	float32
	TextureBody rl.Texture2D
	TextureGun rl.Texture2D
	Frames int
	AniTime float32
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
var GunAngle float64
var dT float32
var RotationSpeed float32
var Car1Wheel rl.Texture2D

const (
	Car1 = iota+1
	Car2 
	Car3
)

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

	Car1Wheel := rl.LoadTexture("./Textures/Car1Wheel.png")
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

	Car1 := Car{Pos: rl.NewVector2(-694, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Nulifier, PowerTime: 10, PowerCooldown: 0.0, Weapon: Minigun, Scale: 5, OffsetSpeed: 1, WheelTexture: Car1Wheel, Frames: 4, Car: Car1}
	Car2 := Car{Pos: rl.NewVector2(-394, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Sat, PowerTime: 10, PowerCooldown: 0.0, Weapon: Minigun, Scale: 5, OffsetSpeed: 1.5}
	Car3 := Car{Pos: rl.NewVector2(-194, World.Y-2200), Texture: Car1_Texture, Health: 600, Power: Emp, PowerTime: 10, PowerCooldown: 0.0, Weapon: Minigun, Scale: 5, OffsetSpeed: 5}

	Gun1 := Gun{PosBody: rl.NewVector2(Car1.Pos.X+300, World.Y-3000), PosGun: rl.NewVector2(Car1.Pos.X+300-(16*3), World.Y-2350), Angle: float32(GunAngle), TextureBody: MiniGunBody, TextureGun: MiniGunTexture, Frames: 0}
	Gun2 := Gun{PosBody: rl.NewVector2(Car2.Pos.X+300, World.Y-3000), PosGun: rl.NewVector2(Car2.Pos.X+300-(16*3), World.Y-2350), Angle: float32(GunAngle), TextureBody: MiniGunBody, TextureGun: MiniGunTexture, Frames: 0}
	Gun3 := Gun{PosBody: rl.NewVector2(Car3.Pos.X+300, World.Y-3000), PosGun: rl.NewVector2(Car3.Pos.X+300-(16*3), World.Y-2350), Angle: float32(GunAngle), TextureBody: MiniGunBody, TextureGun: MiniGunTexture, Frames: 0}
	//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------

	// Game Loop ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	for !rl.WindowShouldClose() {

		dT = rl.GetFrameTime()
		Mouse := rl.GetMousePosition()
		WorldMouse := rl.GetScreenToWorld2D(Mouse, Camera)

		// Position the Guns on the cars ------------------------------
		Gun1 = AttachGunToCar(Gun1, Car1)
		Gun2 = AttachGunToCar(Gun2, Car2)
		Gun3 = AttachGunToCar(Gun3, Car3)
		//-------------------------------------------------------------
		// Gun Following Cursor handler -------------------------------
		Gun1 = GunFollow(WorldMouse, Gun1, Car1)
		Gun2 = GunFollow(WorldMouse, Gun2, Car2)
		Gun3 = GunFollow(WorldMouse, Gun3, Car3)
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
		Car1OffsetX = RandMoveCheckBoundry(0, 900, Car1OffsetX, Car1)
		Car2OffsetX = RandMoveCheckBoundry(1000, 1700, Car2OffsetX, Car2)
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

		
		WheelAnimationPlayer(&Car1, 10, dT, 0)
		rl.DrawTextureEx(Car1.Texture, Car1.Pos, 0.0, Car1.Scale, rl.White) // Render Car 1
		GunAnimationPlay(Car1, &Gun1, 8, dT, 3)
		// rl.DrawTexturePro(Gun1.TextureGun, rl.NewRectangle(0, 0, 39, 7), rl.NewRectangle(Gun1.PosGun.X+85, Gun1.PosGun.Y, 300, 70), rl.NewVector2(85, 30), Gun1.Angle, rl.White)
		rl.DrawTextureEx(Gun1.TextureBody, Gun1.PosBody, 0.0, Car1Minigun.Scale, rl.White)

		rl.DrawTextureEx(Car2.Texture, Car2.Pos, 0.0, Car2.Scale, rl.White)
		GunAnimationPlay(Car2, &Gun2, 8, dT, 3)
		// rl.DrawTexturePro(Gun2.TextureGun, rl.NewRectangle(0, 0, 39, 7), rl.NewRectangle(Gun2.PosGun.X+85, Gun2.PosGun.Y, 300, 70), rl.NewVector2(85, 30), Gun2.Angle, rl.White)
		rl.DrawTextureEx(Gun2.TextureBody, Gun2.PosBody, 0.0, Car1Minigun.Scale, rl.White)

		rl.DrawTextureEx(Car3.Texture, Car3.Pos, 0.0, Car3.Scale, rl.White)
		GunAnimationPlay(Car2, &Gun3, 8, dT, 3)
		// rl.DrawTexturePro(Gun3.TextureGun, rl.NewRectangle(0, 0, 39, 7), rl.NewRectangle(Gun3.PosGun.X+85, Gun3.PosGun.Y, 300, 70), rl.NewVector2(85,30), Gun3.Angle, rl.White)
		rl.DrawTextureEx(Gun3.TextureBody, Gun3.PosBody, 0.0, Car1Minigun.Scale, rl.White)
		
		
		rl.EndMode2D()
		// Camera end ---------------------------------------------------------------------------------------------------------------------

		rl.DrawTextureEx(Car_Hud, rl.NewVector2(0, 0), 0.0, 1, rl.White) // Render Game Hud

		rl.DrawTextEx(rl.GetFontDefault(),fmt.Sprintf("Angle: %0.2f", Gun1.Angle), rl.NewVector2(10, 10), 50, 10, rl.White)

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

// Moves the Gun relative to it's car.
func AttachGunToCar(Gun Gun, Car Car) Gun{
	switch Car.Weapon {
	case Minigun:
		Gun.PosBody = Car.Pos.Add(Car1Minigun.GunBody)
		Gun.PosGun = Car.Pos.Add(Car1Minigun.GunGun)
	}
	return Gun
}
// Makes the Gun follow the Cursor.
func GunFollow(Mouse rl.Vector2, Gun Gun, Car Car) Gun {
		dx := float64(Mouse.X-Gun.PosGun.X)
		dy := float64(Mouse.Y-Gun.PosGun.Y)

		GunAngle := math.Atan2(dy,dx)*180/math.Pi

		DeltaAngle := Gun.Angle - float32(GunAngle)

		switch Car.Weapon {
		case Minigun:
			RotationSpeed = Car1Minigun.RotationSpeed
		}
		
		if DeltaAngle > 0 && Gun.Angle < 45 && Gun.Angle > -135  {
			Gun.Angle -= RotationSpeed
			DeltaAngle += RotationSpeed
		}else if DeltaAngle < 0 && Gun.Angle < 45 && Gun.Angle > -135 {
			Gun.Angle += RotationSpeed
			DeltaAngle -= RotationSpeed
		}else if Gun.Angle >= 45 {
			Gun.Angle = 44.9
		}else if Gun.Angle <= -135 {
			Gun.Angle = -134.9
		}
	return Gun
}

func GunAnimationPlay(Car Car, Gun *Gun, SpriteXSpacing float32, dT float32, Frames int) {
	switch Car.Weapon {
	case Minigun:
		rl.DrawTexturePro(Gun.TextureGun, rl.NewRectangle((42+SpriteXSpacing)* float32(Gun.Frames), 3, 42, 7), rl.NewRectangle(Gun.PosGun.X+85, Gun.PosGun.Y, 323, 70), rl.NewVector2(85, 30), Gun.Angle, rl.White)
	}

	Gun.Frames ++
	if Gun.Frames >= Frames+1 {
		Gun.Frames = 0
	}
}

func WheelAnimationPlayer(Car *Car, SpriteXSpacing float32, dT float32, Frames int) {
	switch Car.Car{
	case Car1:
		rl.DrawRectanglePro(rl.NewRectangle(Car.Pos.X+200, Car.Pos.Y+155, 676, 100), rl.NewVector2(0,0), 0.0, rl.Black)
		rl.DrawTexturePro(Car.WheelTexture, rl.NewRectangle((29+SpriteXSpacing)* float32(Car.Frames), 0, 29, 27), rl.NewRectangle(Car.Pos.X + 214, Car.Pos.Y + 170, 140, 130), rl.NewVector2(0,0), 0.0, rl.White)
		rl.DrawTexturePro(Car.WheelTexture, rl.NewRectangle((29+SpriteXSpacing)* float32(Car.Frames), 0, 29, 27), rl.NewRectangle(Car.Pos.X + 723, Car.Pos.Y + 170, 140, 130), rl.NewVector2(0,0), 0.0, rl.White)
	case Car2:
		rl.DrawTexturePro(Car.WheelTexture, rl.NewRectangle((29+SpriteXSpacing)* float32(Car.Frames),0,0,0), rl.NewRectangle(0,0,0,0), rl.NewVector2(0,0), 0.0, rl.White)
		rl.DrawTexturePro(Car.WheelTexture, rl.NewRectangle((29+SpriteXSpacing)* float32(Car.Frames),0,0,0), rl.NewRectangle(0,0,0,0), rl.NewVector2(0,0), 0.0, rl.White)
	case Car3:
		rl.DrawTexturePro(Car.WheelTexture, rl.NewRectangle((29+SpriteXSpacing)* float32(Car.Frames),0,0,0), rl.NewRectangle(0,0,0,0), rl.NewVector2(0,0), 0.0, rl.White)
		rl.DrawTexturePro(Car.WheelTexture, rl.NewRectangle((29+SpriteXSpacing)* float32(Car.Frames),0,0,0), rl.NewRectangle(0,0,0,0), rl.NewVector2(0,0), 0.0, rl.White)
	}

	Car.Frames ++ 
	if Car.Frames >= Frames+1 {
		Car.Frames = 0
	}
}
