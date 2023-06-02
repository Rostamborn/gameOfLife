package main

import(
    "github.com/eiannone/keyboard"
    "log"
    "os"
)

const(
    UP     = 'k'
	DOWN   = 'j'
	LEFT   = 'h'
	RIGHT  = 'l'
)

type Pointer struct {
    X int
    Y int
}


func HandleInput(pointer *Pointer, cells [][]Cell, ToBeAdded []*Cell, gameState chan bool) {
    err := keyboard.Open()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        _ = keyboard.Close()
    }()
    for {
		key, _, err := keyboard.GetKey()
		if err != nil {
			panic(err) 
        }

		if key == UP {
            pointer.Y--
		} else if key == DOWN {
            pointer.Y++
		} else if key == LEFT {
            pointer.X--
		} else if key == RIGHT {
            pointer.X++
		} else if key == 'p' {
           if !cells[pointer.Y][pointer.X].State {
               ToBeAdded = append(ToBeAdded, &cells[pointer.Y][pointer.X])
               cells[pointer.Y][pointer.X].State = !cells[pointer.Y][pointer.X].State
           }
        } else if key == 's' || key == 'S' {
            gameState <- true
            return
        } else if key == 'q' || key == 'Q' {
            os.Exit(0)
            return
		}
	}
}

func HandleExit() {
    err := keyboard.Open()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        _ = keyboard.Close()
    }()
    for {
		key, _, err := keyboard.GetKey()
		if err != nil {
			panic(err) 
        }
        if key == 'q' || key == 'Q' {
            os.Exit(0)
            return
        }
    }
}
