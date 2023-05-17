package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	lx := make([]int, dx)
	ly := make([]int, dy)
	var picture [][]uint8
	
	for i := range ly {
		var tmp []uint8
		for j := range lx {
			tmp = append(tmp, uint8((i + j) / 2))
		}
		picture = append(picture, tmp)
	}
	
	return picture
}

func main() {
	pic.Show(Pic)
}
