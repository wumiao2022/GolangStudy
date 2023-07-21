package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	randomInt := rand.Intn(100)
	fmt.Println("随机整数:", randomInt)

	// 生成一个随机浮点数
	randomFloat := rand.Float64()
	fmt.Println("随机浮点数:", randomFloat)

	// 生成一个随机布尔值
	randomBool := rand.Intn(2) == 1
	fmt.Println("随机布尔值:", randomBool)

	// 生成一个随机字母
	randomLetter := string(rand.Intn(26) + 97)
	fmt.Println("随机字母:", randomLetter)
}
