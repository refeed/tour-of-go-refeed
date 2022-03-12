package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)


type Image struct{
	width int
	height int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.width, im.height)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x^y), uint8(x^y), 255, 255}
}

func main() {
	m := Image{400, 400}
	pic.ShowImage(m)
}
