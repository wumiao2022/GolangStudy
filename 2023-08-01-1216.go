package main

import "fmt"

func main() {
	// 1. 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 4. 循环打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 定义一个结构体并初始化
	type Person struct {
		name string
		age  int
	}

	person := Person{name: "Alice", age: 25}
	fmt.Println(person)

	// 6. 创建一个切片并添加元素
	numbers := []int{}
	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Println(numbers)

	// 7. 使用指针改变变量的值
	x := 10
	changeValue(&x)
	fmt.Println(x)
}

func changeValue(num *int) {
	*num = 20
}