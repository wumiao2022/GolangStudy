package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    target := rand.Intn(100)
    fmt.Println("Guess the target number between 1 and 100!")
    for guesses := 0; ; guesses++ {
        fmt.Print("Your guess: ")
        var guess int
        _, err := fmt.Scan(&guess)
        if err != nil || guess < 1 || guess > 100 {
            fmt.Println("Invalid input, try again.")
            continue
        }
        if guess < target {
            fmt.Println("Too low, try again.")
        } else if guess > target {
            fmt.Println("Too high, try again.")
        } else {
            fmt.Printf("Congratulations! You guessed it in %d tries.\n", guesses+1)
            break
        }
    }
}