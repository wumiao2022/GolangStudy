package main

import "fmt"

func main() {
	// 1. 将两个整数相加并打印结果
	fmt.Println(2 + 3)

	// 2. 使用for循环打印1到10的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 3. 声明一个整型变量x，并将其初始化为5，然后将x的值乘以2，并将结果打印出来
	x := 5
	x *= 2
	fmt.Println(x)

	// 4. 实现一个函数multiply，接受两个整数参数，返回它们的乘积
	fmt.Println(multiply(3, 4))

	// 5. 声明一个切片slice，其中包含元素1、2、3、4、5，然后使用for循环打印每个元素
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}

	// 6. 声明一个匿名函数，并将其赋值给变量add，使得调用add(2, 3)返回结果5
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))
}

func multiply(a, b int) int {
	return a * b
}