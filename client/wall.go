package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Vector2 struct {
	X, Y float32
}

type Wall struct {
	TopLeft     Vector2
	BottomRight Vector2
	Color       rl.Color
}

func (w *Wall) Draw() {
	rl.DrawRectangle(
		int32(w.TopLeft.X),
		int32(w.TopLeft.Y),
		int32(w.BottomRight.X-w.TopLeft.X),
		int32(w.BottomRight.Y-w.TopLeft.Y),
		w.Color,
	)
}
