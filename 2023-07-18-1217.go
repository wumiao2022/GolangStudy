package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前的日期和时间
	fmt.Println(time.Now())

	// 练习2：计算1到100之间所有数字的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	
	// 练习3：定义一个函数，实现两个整数相加
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(10, 20))

	// 练习4：输出1到10的平方数
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}

	// 练习5：定义一个结构体表示学生，包含姓名和年龄属性
	type Student struct {
		Name string
		Age  int
	}
	student := Student{Name: "Alice", Age: 18}
	fmt.Println(student)
}