package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(100)

	var guess int
	var count int

	for {
		fmt.Print("请猜一个0到100的数字：")
		fmt.Scan(&guess)

		count++

		if guess == num {
			fmt.Printf("恭喜你猜对了！共猜测了%d次。\n", count)
			break
		} else if guess > num {
			fmt.Println("猜大了，请继续猜测。")
		} else {
			fmt.Println("猜小了，请继续猜测。")
		}
	}
}