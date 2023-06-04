package main

import (
	"fmt"
)

const (
	WIDTH       = 80
	HEIGHT      = 30
	BORDER      = "*"
	CELL        = "#"
	EMPTY       = " "
	POINTERICON = "X"
)

func Render(cells [][]Cell, pointer *Pointer) {
	for j := 0; j < HEIGHT; j++ {
		for i := 0; i < WIDTH; i++ {
			if i == 0 || i == WIDTH-1 || j == 0 || j == HEIGHT-1 {
				fmt.Print(BORDER)
			} else if pointer != nil && pointer.X == i && pointer.Y == j {
				fmt.Print(POINTERICON)
			} else if cells[j][i].State {
				fmt.Print(CELL)
			} else {
				fmt.Print(EMPTY)
			}
		}
		fmt.Println()
	}
}
