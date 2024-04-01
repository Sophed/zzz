package main

import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var FPS int
var elements = make(map[string]string)

func drawDebugHud(player *Player) {

	if time.Now().UnixMilli()%500 == 0 {
		FPS = int(rl.GetFPS())
	}

	//elements["FPS"] = strconv.Itoa(FPS)
	elements["Y"] = strconv.FormatFloat(float64(player.Y), 'f', -1, 64)

	offset := int32(20)
	for k, v := range elements {
		rl.DrawText(k+": "+v, 20, offset, 20, rl.White)
		offset += 20
	}

}
