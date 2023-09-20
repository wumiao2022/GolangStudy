package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算1+1，并输出结果
	result := 1 + 1
	fmt.Println(result)

	// 3. 判断一个数是奇数还是偶数，并输出结果
	number := 7
	if number%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 4. 使用循环打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 定义一个切片，包含若干个整数，并使用循环遍历打印出每个元素的值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}