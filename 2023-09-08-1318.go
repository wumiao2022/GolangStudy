package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个随机数
	num := rand.Intn(100)

	// 判断随机数的奇偶性
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}