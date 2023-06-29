package main

import "fmt"

func main() {
	// 1. 输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 定义并输出一个整数变量
	num := 10
	fmt.Println(num)

	// 3. 计算两个整数的和并输出结果
	a := 1
	b := 2
	sum := a + b
	fmt.Println(sum)

	// 4. 判断一个数是否为偶数，并输出结果
	num2 := 6
	if num2%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 5. 输出0-9的所有整数
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 6. 输出一个字符串的长度
	str := "Golang"
	fmt.Println(len(str))

	// 7. 定义并输出一个切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 8. 定义并输出一个结构体
	type person struct {
		name string
		age  int
	}
	p1 := person{name: "Alice", age: 20}
	fmt.Println(p1)

	// 9. 调用一个函数并输出结果
	result := add(3, 4)
	fmt.Println(result)
}

// 定义一个函数用于计算两个整数的和
func add(a, b int) int {
	return a + b
}