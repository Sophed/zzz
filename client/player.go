package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X, Y    int32
	Texture rl.Texture2D
}

func NewPlayer() *Player {

	texture := rl.LoadTexture(ASSETS_DIR + "player.png")
	return &Player{
		WINDOW_WIDTH / 2,
		WINDOW_HEIGHT / 2,
		texture,
	}

}

func (p *Player) GetHitbox() Hitbox {
	return Hitbox{
		Vector2{
			p.X,
			p.Y,
		},
		Vector2{
			p.X + p.Texture.Width,
			p.Y + p.Texture.Height,
		},
	}
}

func (p *Player) Draw() {
	texture := p.Texture
	texture.Height = texture.Height * PIXEL_SCALE
	texture.Width = texture.Width * PIXEL_SCALE
	rl.DrawTexture(
		texture,
		p.X,
		p.Y,
		rl.White,
	)
}
