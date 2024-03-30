package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Vector2 struct {
	X, Y int32
}

type Wall struct {
	TopLeft     Vector2
	BottomRight Vector2
	Color       rl.Color
}

func (w *Wall) Draw() {
	rl.DrawRectangle(
		w.TopLeft.X,
		w.TopLeft.Y,
		w.BottomRight.X-w.TopLeft.X,
		w.BottomRight.Y-w.TopLeft.Y,
		w.Color,
	)
}

func (w *Wall) Colliding(v Vector2) bool {
	if v.X < w.TopLeft.X && v.X > w.BottomRight.X {
		return true
	}
	if v.Y < w.TopLeft.Y && v.Y > w.BottomRight.Y {
		return true
	}
	return false
}
