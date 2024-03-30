package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Hitbox struct {
	TopLeft     Vector2
	BottomRight Vector2
}

func (h *Hitbox) Width() int32 {
	return h.BottomRight.X - h.TopLeft.X
}

func (h *Hitbox) Height() int32 {
	return h.BottomRight.Y - h.TopLeft.Y
}

func (h *Hitbox) Draw() {
	rl.DrawRectangleLines(
		h.TopLeft.X,
		h.TopLeft.Y,
		h.Width()*PIXEL_SCALE,
		h.Height()*PIXEL_SCALE,
		rl.Red,
	)
}
