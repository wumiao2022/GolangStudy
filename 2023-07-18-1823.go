package main
  
import (
	"fmt"
	"math/rand"
	"time"
)
  
func main() {
	rand.Seed(time.Now().UnixNano())
  
	for i := 0; i < 5; i++ {
		num := rand.Intn(10)
  
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}