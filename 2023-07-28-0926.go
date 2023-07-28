package main

import (
	"fmt"
)

func main() {
	// 练习：使用循环打印出1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习：使用数组存储并打印出5个人的姓名
	names := [5]string{"Tom", "Alice", "John", "Emily", "Mike"}
	for _, name := range names {
		fmt.Println(name)
	}

	// 练习：使用切片存储并打印出5个人的年龄
	ages := []int{25, 30, 28, 35, 27}
	for _, age := range ages {
		fmt.Println(age)
	}
}

// 这是一个多练习的示例代码，包含了使用循环打印整数、计算和、打印数组和切片等基本操作。请运行代码以查看结果。