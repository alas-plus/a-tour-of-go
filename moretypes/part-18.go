package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	ay := make([][]uint8, dy)
	for y := range ay {
		ax := make([]uint8, dx)
		for x := range ax {
			val := uint8(x ^ y) //x^y, (x+y)/2, x*y
			ax[x] = val
		}
		ay[y] = ax
	}
	return ay
}

func main() {
	pic.Show(Pic)
}
