package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X, Y     float32
	Velocity Vector2
	Texture  rl.Texture2D
}

func NewPlayer() *Player {

	texture := rl.LoadTexture(ASSETS_DIR + "player.png")
	return &Player{
		WINDOW_WIDTH / 2,
		WINDOW_HEIGHT / 2,
		Vector2{
			0,
			0,
		},
		texture,
	}

}

const PLAYER_MAX_SPEED = 500
const PLAYER_ACCELERATION = 3

func (p *Player) Move() {
	p.X += p.Velocity.X * rl.GetFrameTime()
	p.Y += p.Velocity.Y * rl.GetFrameTime()
}

func (p *Player) HandleInput() {

	if rl.IsKeyDown(rl.KeyA) {
		if !(math.Abs(float64(p.Velocity.X))-PLAYER_ACCELERATION > PLAYER_MAX_SPEED) {
			p.Velocity.X -= PLAYER_ACCELERATION
		}
	}
	if rl.IsKeyDown(rl.KeyD) {
		if !(math.Abs(float64(p.Velocity.X))+PLAYER_ACCELERATION > PLAYER_MAX_SPEED) {
			p.Velocity.X += PLAYER_ACCELERATION
		}
	}

	if p.Velocity.X > 0 {
		p.Velocity.X--
	} else if p.Velocity.X < 0 {
		p.Velocity.X++
	}

}

func (p *Player) GetHitbox() Hitbox {
	return Hitbox{
		Vector2{
			p.X,
			p.Y,
		},
		Vector2{
			p.X + float32(p.Texture.Width),
			p.Y + float32(p.Texture.Height),
		},
	}
}

func (p *Player) Draw() {
	texture := p.Texture
	texture.Height = texture.Height * PIXEL_SCALE
	texture.Width = texture.Width * PIXEL_SCALE
	rl.DrawTexture(
		texture,
		int32(p.X),
		int32(p.Y),
		rl.White,
	)
}
