package main

import "fmt"

func main() {
	// 示例1：输出Hello World
	fmt.Println("Hello World")

	// 示例2：计算两数之和
	a := 5
	b := 10
	sum := a + b
	fmt.Println("Sum:", sum)

	// 示例3：循环输出1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 示例4：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 示例5：使用数组存储一组数字，并计算它们的平均值
	numbers := [5]int{1, 2, 3, 4, 5}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	fmt.Println("Average:", average)
}