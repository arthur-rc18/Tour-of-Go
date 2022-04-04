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

	// The NewRGBA function returns a new RGBA image with the given bounds
	m := image.NewRGBA(image.Rect(0, 0, 100, 100)) // Rect is shorthand for rectangle{Pt(x0,y0),Pt(x1,y1)}
	fmt.Println(m.Bounds())                        // The function Bounds
	fmt.Println(m.At(0, 0).RGBA())                 // The At function explained previously
}
