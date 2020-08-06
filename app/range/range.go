package main

import "golang.org/x/tour/pic"

func Pic(x, y int) [][]uint8 {
	image := make([][]uint8, y)
	for i := 0; i < y; i++ {
		image[i] = make([]uint8, x)
		for k := 0; k < x; k++ {
			image[i][k] = uint8(k * i)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
