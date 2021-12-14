package internal

import (
	"github.com/faiface/pixel"
)

type World struct {
	width  float64
	height float64

	bg pixel.Picture
}

func NewWorld(width, height float64) *World {
	return &World{
		width:  width,
		height: height,
	}

}
func (w *World) AddBackground (path string) error {
	bg, err := loadPicture(path)
	if err != nil {
		return err
	}
	w.bg = bg
	return nil
}
func (w World) Draw(t pixel.Target){
	spriteBg := pixel.NewSprite(w.bg, w.bg.Bounds())
	bgBatch := pixel.NewBatch(&pixel.TrianglesData{}, w.bg)
	bgBatch.Clear()

	for x := 0.0; x <= w.width; += spriteBg.Frame().W(){
		for y := 0.0; y <= w.height; += spriteBg.Frame().H(){
			pos := pixel.V(x, y)
			spriteBg.Draw(bgBatch, pixel.IM.Moved(pos))

		}
	}
	bgBatch.Draw(t)
}
