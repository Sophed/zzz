package main

import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ASSETS_DIR = "assets/"
const WINDOW_WIDTH = 1280
const WINDOW_HEIGHT = 720
const PIXEL_SCALE = 4

const BOX_PADDING = PIXEL_SCALE * 16
const BOX_WIDTH = PIXEL_SCALE * 4

var FPS int

func main() {

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "zzz")
	defer rl.CloseWindow()

	setBox()

	player := NewPlayer()
	//player.Velocity.X = 20

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		player.HandleInput()
		player.Move()
		player.Gravity()
		player.Draw()
		hitbox := player.GetHitbox()
		hitbox.Draw()

		drawBox()
		drawFPS()
		rl.EndDrawing()
	}

}

func drawFPS() {
	if time.Now().UnixMilli()%500 == 0 {
		FPS = int(rl.GetFPS())
	}
	rl.DrawText("FPS: "+strconv.Itoa(FPS), 20, 20, 20, rl.RayWhite)
}
