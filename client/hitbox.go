package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Hitbox struct {
	TopLeft     Vector2
	BottomRight Vector2
}

func (h *Hitbox) Width() float32 {
	return h.BottomRight.X - h.TopLeft.X
}

func (h *Hitbox) Height() float32 {
	return h.BottomRight.Y - h.TopLeft.Y
}

func (h *Hitbox) Draw() {
	rl.DrawRectangleLines(
		int32(h.TopLeft.X),
		int32(h.TopLeft.Y),
		int32(h.Width()),
		int32(h.Height()),
		rl.Red,
	)
}
