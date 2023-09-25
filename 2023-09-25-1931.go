package main

import "fmt"

func main() {
	// 实例1：打印 Hello World
	fmt.Println("Hello World")
	
	// 实例2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 实例3：判断一个数是否为偶数
	num := 5
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	
	// 实例4：遍历数组
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
	
	// 实例5：使用函数计算阶乘
	num := 5
	factorial := calcFactorial(num)
	fmt.Println(num, "factorial:", factorial)
}

func calcFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calcFactorial(n-1)
}
