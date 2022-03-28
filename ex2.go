package main

func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for z := 0; z < dx; z++ {
			s[z] = uint8((y + z) / 2)
		}
		x[y] = s
	}
	return x
}

func main() {
	pic.show(Pic)
}
