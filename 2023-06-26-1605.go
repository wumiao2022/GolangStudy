package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100)
    var guess int
    fmt.Println("Guess the number between 0 and 100")
    for {
        fmt.Print("Your guess: ")
        fmt.Scan(&guess)
        if guess < target {
            fmt.Println("Too low")
        } else if guess > target {
            fmt.Println("Too high")
        } else {
            fmt.Println("Congratulations, you guessed the number!")
            break
        }
    }
}