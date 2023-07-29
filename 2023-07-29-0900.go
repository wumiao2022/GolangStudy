package main

import "fmt"

func main() {
	// 练习1: 输出Hello World
	fmt.Println("Hello World")

	// 练习2: 求两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数
	num := 24
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 打印1-10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 使用切片获取指定范围内的元素
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	startIndex := 2
	endIndex := 7
	slicedNums := nums[startIndex:endIndex+1]
	fmt.Println(slicedNums)
}