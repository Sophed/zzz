package main

import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var FPS int
var elements = make(map[string]string)

func drawDebugHud() {

	if time.Now().UnixMilli()%500 == 0 {
		FPS = int(rl.GetFPS())
	}

	elements["FPS"] = strconv.Itoa(FPS)

	offset := int32(20)
	for k, v := range elements {
		rl.DrawText(k+": "+v, 20, offset, 20, rl.White)
		offset += 20
	}

}