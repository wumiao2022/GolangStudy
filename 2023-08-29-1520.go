package main

import "fmt"

func main() {
	// 示例1：打印Hello World
	fmt.Println("Hello World!")

	// 示例2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 示例3：判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 示例4：遍历数组打印元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 示例5：定义和使用结构体
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 25}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}