package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic = [][]uint8{}

	for y := 0; y < dy; y++ {
		pic = append(pic, []uint8{})
		for x :=0; x < dx; x++ {
			pic[y] = append(pic[y], uint8(x^y))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
