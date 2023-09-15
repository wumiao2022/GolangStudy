package main

import "fmt"

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：求两个数的和
	x := 5
	y := 10
	sum := x + y
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用数组存储五个整数，并计算它们的和
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)
}