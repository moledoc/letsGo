package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	imageFunc := func(x, y int) uint8 {
		return uint8(x * y)
	}
	val := imageFunc(x, y)
	return color.RGBA{val, val, 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
