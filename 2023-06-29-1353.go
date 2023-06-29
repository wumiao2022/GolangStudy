package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个随机数，并打印该随机数
	randomNum := rand.Intn(100)
	fmt.Println(randomNum)

	// 练习2: 生成一个随机数切片，包含10个随机数，并打印该切片
	var randomSlice []int
	for i := 0; i < 10; i++ {
		randomSlice = append(randomSlice, rand.Intn(100))
	}
	fmt.Println(randomSlice)

	// 练习3: 计算两个整数的和并返回
	num1 := 20
	num2 := 30
	sum := num1 + num2
	fmt.Println(sum)

	// 练习4: 判断一个数是否为偶数，如果是返回true，否则返回false
	checkEven := func(num int) bool {
		return num%2 == 0
	}

	fmt.Println(checkEven(10))
	fmt.Println(checkEven(15))
}
