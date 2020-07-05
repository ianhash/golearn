package exercise

import (
	"image"
	"image/color"
)

// Image my image
type Image struct {
	Width, Height int
}

// ColorModel color model of my image
func (it *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds get bounds of my image
func (it *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, it.Width, it.Height)
}

// At get color at (x,y)
func (it *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(x * x), uint8(y * y)}
}
