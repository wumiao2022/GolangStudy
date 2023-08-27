package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个随机数
	num := rand.Intn(100)
	fmt.Println("随机数:", num)

	// 判断随机数的奇偶性
	if num%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}

	// 判断随机数的大小
	if num > 50 {
		fmt.Println("随机数大于50")
	} else {
		fmt.Println("随机数小于等于50")
	}
}
