package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X, Y     float32
	Velocity Vector2
	Texture  rl.Texture2D
	Jumping  bool
}

func NewPlayer() *Player {

	texture := rl.LoadTexture(ASSETS_DIR + "player.png")
	texture.Height = texture.Height * PIXEL_SCALE
	texture.Width = texture.Width * PIXEL_SCALE

	return &Player{
		750,
		200,
		Vector2{
			0,
			0,
		},
		texture,
		false,
	}

}

const PLAYER_MAX_SPEED = 500
const PLAYER_ACCELERATION = 3
const PLAYER_GRAVITY = 1
const PLAYER_JUMP = 60

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
	if rl.IsKeyDown(rl.KeySpace) {
		p.Jump()
	}

	if p.Velocity.X > 0 {
		p.Velocity.X--
	} else if p.Velocity.X < 0 {
		p.Velocity.X++
	}

}

func (p *Player) Jump() {
	if !p.OnGround() || p.Jumping {
		return
	}
	p.Jumping = true
	p.Velocity.Y -= PLAYER_JUMP * 10
}

func (p *Player) Gravity() {
	p.Velocity.Y += PLAYER_GRAVITY
	p.OnGround()
}

func (p *Player) OnGround() bool {
	if p.GetHitbox().BottomRight.Y > float32(getFloor()-1) {
		p.Jumping = false
		p.Velocity.Y = 0
		return true
	}
	return false
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
	rl.DrawTexture(
		p.Texture,
		int32(p.X),
		int32(p.Y),
		rl.White,
	)
}
