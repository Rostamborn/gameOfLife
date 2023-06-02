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
    // we check for the changes for the next iteration
    // then at the end apply the changes
    ToDie := make([]*Cell, 0)
    ToBeBorn := make([]*Cell, 0)
    for j := 0; j < HEIGHT; j++ {
        for i := 0; i < WIDTH; i++ {
            cells[j][i].CheckForNeighbours(cells)
            // logic of game of life
            if cells[j][i].State {
                if len(cells[j][i].LiveNeighbours) < 2 || 3 < len(cells[j][i].LiveNeighbours) {
                    ToDie = append(ToDie, &cells[j][i])
                }
            } else if !cells[j][i].State {
                if len(cells[j][i].LiveNeighbours) == 3 {
                    ToBeBorn = append(ToBeBorn, &cells[j][i])
                }
            }
        }
    }
    for _, v := range ToBeBorn {
        v.BeBorn()
    }
    for _, v := range ToDie {
        v.Die()
    }
}

func main() {
    pointer := Pointer{X: WIDTH/2, Y: HEIGHT/2}
    gameState := make(chan bool)
    cells := make([][]Cell, HEIGHT)
    ToBeAdded := make([]*Cell, 0)
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
    // cells[11][10].State = true
    // cells[12][10].State = true
    // cells[11][11].State = true
    // cells[12][9].State = true
    // cells[22][14].State = true
    // cells[12][12].State = true
    // cells[12][13].State = true
    // cells[12][14].State = true
    // cells[12][15].State = true
    // cells[12][16].State = true
    // cells[12][17].State = true
    // cells[12][18].State = true
    // cells[12][19].State = true
    

    ticker := time.NewTicker(200 * time.Millisecond)
    defer ticker.Stop()

    go HandleInput(&pointer, cells, ToBeAdded, gameState)

    OuterLoop:
    for range ticker.C {
        select{
        case <- gameState:
            break OuterLoop
        default:
            clearScreen()
            Render(cells, &pointer)
        }
    }
    ticker.Stop()

    for _, v := range ToBeAdded {
        v.State = true
    }

    newTicker := time.NewTicker(500 * time.Millisecond)
    defer newTicker.Stop()
    go HandleExit()

    for range newTicker.C {
        select{
        default:
            clearScreen()
            Render(cells, nil)
            updateNeighbours(cells)
        }
    }
}
