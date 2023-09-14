package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")

	// 练习1：打印当前时间
	fmt.Println("Current time:", time.Now())

	// 练习2：计算1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 100:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：找出一个数组中的最大值
	numbers := []int{34, 56, 12, 87, 35, 9}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max number in the array:", max)
}