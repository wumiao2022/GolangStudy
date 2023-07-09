package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成一个随机整数并判断其是否为偶数
	num := rand.Intn(100)
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习2：从1到10循环打印数字，并打印出每个数字的平方
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d 的平方是 %d\n", i, i*i)
	}

	// 练习3：计算并打印斐波那契数列的前20个数字
	fmt.Println("斐波那契数列的前20个数字：")
	fibonacci := []int{0, 1}
	for i := 2; i < 20; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)
}
```