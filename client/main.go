package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const DEBUG = true
const ASSETS_DIR = "assets/"
const WINDOW_WIDTH = 1280
const WINDOW_HEIGHT = 720
const PIXEL_SCALE = 4

const BOX_PADDING = PIXEL_SCALE * 16
const BOX_WIDTH = PIXEL_SCALE * 4

func main() {

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "zzz")
	defer rl.CloseWindow()

	setBox()

	player := NewPlayer()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		player.HandleInput()
		player.Move()
		player.Gravity()
		player.Draw()
		hitbox := player.GetHitbox()

		drawBox()

		if DEBUG {
			drawDebugHud(player)
			player.DrawDirection()
			hitbox.Draw()
		}

		rl.EndDrawing()
	}

}
