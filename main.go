package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
    cells := make([][]Cell, WIDTH)
    for i := range cells {
        cells[i] = make([]Cell, HEIGHT)
    }
    fmt.Println("Hello, World!")
    ticker := time.NewTicker(200 * time.Millisecond)
    defer ticker.Stop()

    for range ticker.C {
        select{
        default:
            clearScreen()
            Render(cells)
        }
    }
}
