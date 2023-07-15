package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 练习1：打印1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：生成并打印一个随机的整数
	randomNum := rand.Intn(100)
	fmt.Println(randomNum)

	// 练习3：计算两个整数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 练习4：判断一个整数是否为偶数，并打印结果
	number := 7
	if number%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

注意：以上代码仅供参考，可能需要根据具体环境和需求进行调整。