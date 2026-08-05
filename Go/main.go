package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetWindowState(rl.FlagFullscreen,rl.FlagMaximized)
	rl.InitWindow(0,0)
}