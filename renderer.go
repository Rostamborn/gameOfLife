package main

import(
    "fmt"
)

const(
    WIDTH = 80
    HEIGHT = 30
    BORDER = "#"
    CELL = "*"
    EMPTY = " "
)

func Render(cells [][]Cell) {
    for j := 0; j < HEIGHT; j++ {
        for i := 0; i < WIDTH; i++ {
            if i == 0 || i == WIDTH - 1 || j == 0 || j == HEIGHT - 1 {
                fmt.Print(BORDER)
            } else if cells[i][j].State {
                fmt.Print(CELL)
                cells[i][j].CheckForNeighbours(cells)
                cells[i][j].CheckState()
            } else {
                fmt.Print(EMPTY)
            }
        }
        fmt.Println()
    }
}
