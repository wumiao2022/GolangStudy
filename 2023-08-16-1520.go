package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：使用for循环打印从1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 练习2：使用for循环打印从10到1的整数
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 练习3：使用for循环打印0到100之间所有能被3整除的数
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// 练习4：使用for循环计算1到100所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Sum: %d\n", sum)

	// 练习5：使用for循环计算1到100所有偶数的和
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of even numbers: %d\n", sum)

	// 练习6：使用for循环计算1到100所有奇数的和
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of odd numbers: %d\n", sum)

	// 练习7：使用for循环打印当前时间，每隔1秒更新一次
	for {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}