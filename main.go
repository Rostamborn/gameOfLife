package main

import (
	"os"
	"os/exec"
	"time"
)


func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func updateNeighbours(cells [][]Cell) {
    for j := 0; j < HEIGHT; j++ {
        for i := 0; i < WIDTH; i++ {
            cells[j][i].CheckForNeighbours(cells)
            cells[j][i].CheckState()
        }
    }
}

func main() {
    cells := make([][]Cell, HEIGHT)
    for i := range cells {
        cells[i] = make([]Cell, WIDTH)
    }
    for j := 0; j < HEIGHT; j++ {
        for i := 0; i < WIDTH; i++ {
            cells[j][i] = Cell{X: i, Y: j, State: false}
        }
    }
    // cells[10][10].State = true
    // cells[10][11].State = true
    // cells[9][12].State = true
    cells[11][10].State = true
    cells[12][10].State = true
    cells[11][11].State = true
    cells[12][9].State = true
    // cells[22][14].State = true
    // cells[12][12].State = true
    // cells[12][13].State = true
    // cells[12][14].State = true
    // cells[12][15].State = true
    // cells[12][16].State = true
    // cells[12][17].State = true
    // cells[12][18].State = true
    // cells[12][19].State = true
    

    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        select{
        default:
            clearScreen()
            Render(cells)
            updateNeighbours(cells)
        }
    }
}
