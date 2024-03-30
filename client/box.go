package main

import rl "github.com/gen2brain/raylib-go/raylib"

var left_wall Wall
var right_wall Wall
var top_wall Wall
var bottom_wall Wall

func drawBox() {
	left_wall.Draw()
	right_wall.Draw()
	top_wall.Draw()
	bottom_wall.Draw()
}

func getFloor() int32 {
	return (WINDOW_HEIGHT - BOX_PADDING) - BOX_WIDTH
}

func setBox() {
	left_wall = Wall{
		Vector2{
			BOX_PADDING,
			BOX_PADDING,
		},
		Vector2{
			BOX_PADDING + BOX_WIDTH,
			WINDOW_HEIGHT - BOX_PADDING,
		},
		rl.LightGray,
	}

	right_wall = Wall{
		Vector2{
			(WINDOW_WIDTH - BOX_PADDING) - BOX_WIDTH,
			BOX_PADDING,
		},
		Vector2{
			WINDOW_WIDTH - BOX_PADDING,
			WINDOW_HEIGHT - BOX_PADDING,
		},
		rl.LightGray,
	}

	top_wall = Wall{
		Vector2{
			BOX_PADDING,
			BOX_PADDING,
		},
		Vector2{
			WINDOW_WIDTH - BOX_PADDING,
			BOX_PADDING + BOX_WIDTH,
		},
		rl.LightGray,
	}

	bottom_wall = Wall{
		Vector2{
			BOX_PADDING,
			(WINDOW_HEIGHT - BOX_PADDING) - BOX_WIDTH,
		},
		Vector2{
			WINDOW_WIDTH - BOX_PADDING,
			WINDOW_HEIGHT - BOX_PADDING,
		},
		rl.LightGray,
	}

}
