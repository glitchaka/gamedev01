package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	pixelgl.Run(run)
}
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Friends of GO: Space Game",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}
	win.Clear(colornames.Aqua)
	for !win.Closed() {
		win.Update()
	}
	world := gamedev01.NewWorld(windowWidth, windowHeight)
	if err := world.AddBackground("resources/background.png"); err != nil {
		log.Fatal(err)
	}
	world.Draw(win)
	for !win.Closed() {
		win.Update()
	}
}
