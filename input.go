package main

import(
    "github.com/eiannone/keyboard"
    "log"
)

func HandleInput() {
    err := keyboard.Open()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        _ = keyboard.Close()
    }()


}
