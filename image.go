package main

import (
	"fmt"
	"image"
)

// Package image defines the IMage interface:
// package image

// type Image interface{
//	ColorModel() color.Model	ColorModel returns the Image's color model.
//	Bounds() Rectangle	// Bounds returns the domain for which At can return non-zero color.
//	At(x, y int) color.Color	At returns the color of the pixel at (x,y)
//}

// This is how is structured the Image interface

func main() {

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
