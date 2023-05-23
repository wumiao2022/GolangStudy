package main

import (
	"fmt"
)

func main() {
	// 练习1：输出0到9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 练习2：输出1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：输出10到1的倒序
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	// 练习4：输出100以内的偶数
	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5：输出质数
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}