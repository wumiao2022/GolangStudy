package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	fmt.Println("请猜一个0到100之间的整数")
	for i := 1; i <= 10; i++ {
		fmt.Printf("第%d次猜测：", i)
		var guess int
		fmt.Scan(&guess)
		if guess == num {
			fmt.Println("恭喜你猜对了!")
			return
		} else if guess > num {
			fmt.Println("猜大了")
		} else {
			fmt.Println("猜小了")
		}
	}
	fmt.Println("很遗憾，没有猜中。正确答案是", num)
}