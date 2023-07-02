package main

import "fmt"

func main() {
    // 实例1: 打印 "Hello World!"
	fmt.Println("Hello World!")

	// 实例2: 计算两个整数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3: 判断一个数是否为偶数并打印结果
	num := 9
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 实例4: 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 实例5: 定义一个切片并遍历打印
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}