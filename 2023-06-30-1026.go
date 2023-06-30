package main

import (
	"fmt"
)

func main() {
	// 每日多练，学习Golang

	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算两个整数的和
	a := 5
	b := 10
	sum := a + b
	fmt.Println("Sum =", sum)

	// 3. 判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 4. 使用循环计算1到100的和
	sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 100 =", sum)

	// 5. 定义一个结构体并使用结构体字段
	type Person struct {
		Name string
		Age  int
	}
	person := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println("Person:", person)

	// 6. 使用切片进行数组元素的删除和插入操作
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original numbers:", numbers)
	// 删除元素
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println("After deleting element at index", index, ", numbers:", numbers)
	// 插入元素
	index = 1
	numbers = append(numbers[:index], append([]int{10}, numbers[index:]...)...)
	fmt.Println("After inserting 10 at index", index, ", numbers:", numbers)

	// 7. 使用map存储键值对并进行遍历
	persons := map[string]int{
		"John": 30,
		"Jane": 25,
		"Bob":  35,
	}
	fmt.Println("Persons:")
	for name, age := range persons {
		fmt.Println(name, "is", age, "years old")
	}

	// 8. 使用递归函数计算阶乘
	fmt.Println("Factorial of 5 is", factorial(5))

	// 9. 使用defer关键字延迟执行函数
	defer fmt.Println("Deferred statement")
	fmt.Println("Regular statement")

	// 10. 使用panic和recover进行错误处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from error:", err)
		}
	}()
	panic("Error occurred")

}

// 阶乘函数
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}