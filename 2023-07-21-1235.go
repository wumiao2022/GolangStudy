package main

import "fmt"

func main() {
	// 1. 输出Hello, World!到控制台
	fmt.Println("Hello, World!")

	// 2. 定义和输出一个整数变量
	num := 10
	fmt.Println(num)

	// 3. 计算并输出两个整数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Println(sum)

	// 4. 判断一个数是否为偶数，并输出结果
	c := 7
	if c%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 5. 使用for循环输出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 6. 创建一个切片，并输出其中的元素
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 7. 使用函数递归计算阶乘
	n := 5
	factorial := calculateFactorial(n)
	fmt.Println(factorial)
}

func calculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorial(n-1)
}