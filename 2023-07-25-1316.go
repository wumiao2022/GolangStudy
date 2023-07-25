package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成1到100之间的随机数，并判断是否为偶数
	num := rand.Intn(100) + 1
	if num%2 == 0 {
		fmt.Printf("%d 是一个偶数\n", num)
	} else {
		fmt.Printf("%d 是一个奇数\n", num)
	}

	// 练习2：使用for循环输出1到10之间的奇数
	fmt.Println("1到10之间的奇数：")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习3：计算1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Printf("1到100之间所有偶数的和为：%d\n", sum)
}