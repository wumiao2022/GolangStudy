package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	num := rand.Intn(100)

	// 输出随机数
	fmt.Println(num)

	// 判断随机数是否为偶数
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}