package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	myslice := make([][]uint8, dy)
	for i := range myslice {
		myslice[i] = make([]uint8, dx)
		for k := range myslice[i] {
			myslice[i][k] = uint8((k + i) ^ 2/2)
		}
	}
	return myslice
}

func main() {
	pic.Show(Pic)
}
