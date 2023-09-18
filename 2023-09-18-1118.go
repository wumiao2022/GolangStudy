package main

import "fmt"

func main() {
	// 练习1: 输出Hello World
	fmt.Println("Hello World")

	// 练习2: 定义和打印变量
	var name string = "Alice"
	age := 20
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	// 练习3: 数组的遍历和求和
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 练习4: 切片的基本操作
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Fruits:", fruits)
	fruits = append(fruits, "grape")
	fmt.Println("New Fruits:", fruits[1:3])

	// 练习5: 使用函数计算阶乘
	factorial := func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}
	fmt.Println("Factorial of 5:", factorial(5))
}