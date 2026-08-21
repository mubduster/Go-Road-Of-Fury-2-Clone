package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GunPos struct {
	GunBody       rl.Vector2
	GunGun        rl.Vector2
	Scale         float32
	RotationSpeed float32
}

const (
	Minigun = iota
	PulseGun
	Laser
)

var Car1Minigun GunPos = GunPos{
	GunBody:       rl.NewVector2(405, -90),
	GunGun:        rl.NewVector2(375, -120),
	Scale:         7,
	RotationSpeed: 1,
}
