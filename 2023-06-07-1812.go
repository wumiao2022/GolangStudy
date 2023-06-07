package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 生成一个0到100之间的随机数
	target := rand.Intn(100)

	// 打印提示信息
	fmt.Println("Guess the number! It is between 0 and 100.")

	// 循环猜测直到猜中为止
	for {
		// 接收用户输入的猜测数
		var guess int
		fmt.Scan(&guess)

		// 判断是否猜中了
		if guess == target {
			fmt.Println("Congratulations! You got it!")
			break
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Oops. Your guess was HIGH.")
		}
	}
}