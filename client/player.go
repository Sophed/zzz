package main

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X, Y      float32
	Velocity  Vector2
	Texture   rl.Texture2D
	Jumping   bool
	Dashing   bool
	CanDash   bool
	LastDash  int64
	Direction int
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
		false,
		true,
		time.Now().UnixMilli(),
		1,
	}

}

const PLAYER_MAX_SPEED = 500
const PLAYER_ACCELERATION = 3
const PLAYER_GRAVITY = 1
const PLAYER_JUMP = 60
const PLAYER_DASH_STENGTH = 800
const PLAYER_DASH_DURATION = 200

func (p *Player) Move() {
	p.X += p.Velocity.X * rl.GetFrameTime()
	p.Y += p.Velocity.Y * rl.GetFrameTime()
}

func (p *Player) HandleInput() {

	if rl.IsKeyDown(rl.KeyA) && !p.Dashing {
		if !(math.Abs(float64(p.Velocity.X))-PLAYER_ACCELERATION > PLAYER_MAX_SPEED) {
			p.Velocity.X -= PLAYER_ACCELERATION
		}
		if !rl.IsKeyDown(rl.KeyD) {
			p.Direction = -1
		}
	}
	if rl.IsKeyDown(rl.KeyD) && !p.Dashing {
		if !(math.Abs(float64(p.Velocity.X))+PLAYER_ACCELERATION > PLAYER_MAX_SPEED) {
			p.Velocity.X += PLAYER_ACCELERATION
		}
		if !rl.IsKeyDown(rl.KeyA) {
			p.Direction = 1
		}
	}
	if rl.IsKeyDown(rl.KeySpace) {
		p.Jump()
	}
	if (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsMouseButtonDown(rl.MouseButtonLeft)) && !p.Dashing && p.CanDash && !p.OnGround() {
		p.Dash()
	}

	if time.Now().UnixMilli()-p.LastDash >= PLAYER_DASH_DURATION {
		p.Dashing = false
	}

	if p.Velocity.X > 0 {
		p.Velocity.X--
	} else if p.Velocity.X < 0 {
		p.Velocity.X++
	}

}

func (p *Player) Dash() {
	p.Dashing = true
	p.CanDash = false
	p.LastDash = time.Now().UnixMilli()
	p.Velocity.Y = 0
	p.ClipSpeed()
	dx := 0
	//dy := 0
	if rl.IsKeyDown(rl.KeyA) {
		dx -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		dx += 1
	}
	if dx == 0 {
		dx = p.Direction
	}
	/*if rl.IsKeyDown(rl.KeyW) {
		dy -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		dy += 1
	}*/
	p.Velocity.X += float32(dx) * (PLAYER_DASH_STENGTH)
	//p.Velocity.Y += float32(dy) * (PLAYER_DASH_STENGTH * 0.6)
}

func (p *Player) ClipSpeed() {
	p.Velocity.X = PLAYER_MAX_SPEED * float32(p.Direction)
}

func (p *Player) Jump() {
	if !p.OnGround() || p.Jumping {
		return
	}
	p.Jumping = true
	p.Velocity.Y -= PLAYER_JUMP * 10
}

func (p *Player) Gravity() {
	p.OnGround()
	if p.Dashing {
		return
	}
	p.Velocity.Y += PLAYER_GRAVITY
}

func (p *Player) OnGround() bool {
	if p.GetHitbox().BottomRight.Y > float32(getFloor()-1) {
		p.Jumping = false
		p.CanDash = true
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

const DIRECTION_LINE_LENGTH = 100

func (p *Player) DrawDirection() {
	center := Vector2{
		p.X + (float32(p.Texture.Width) / 2),
		p.Y + (float32(p.Texture.Height) / 2),
	}
	rl.DrawLine(
		int32(center.X),
		int32(center.Y),
		int32(center.X)+(DIRECTION_LINE_LENGTH*int32(p.Direction)),
		int32(center.Y),
		rl.Green,
	)
}
