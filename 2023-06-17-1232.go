package main

import (
	"fmt"
)

func main() {
	// 练习1：判断奇偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习2：for循环打印1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：数组遍历求和
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)

	// 练习4：判断素数
	num2 := 7
	flag := true
	for i := 2; i < num2; i++ {
		if num2%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("是素数")
	} else {
		fmt.Println("不是素数")
	}
}