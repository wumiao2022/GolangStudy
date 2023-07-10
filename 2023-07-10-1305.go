package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并打印结果
	num1 := 5
	num2 := 3
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：在数组中查找某个特定的元素并打印其索引
	arr := []int{2, 4, 6, 8, 10}
	target := 6
	for i, v := range arr {
		if v == target {
			fmt.Println("Index of", target, "is", i)
			break
		}
	}

	// 练习5：使用for循环打印1到10的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}