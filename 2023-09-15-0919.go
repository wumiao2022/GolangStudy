package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算并打印1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 10:", sum)

	// 3. 输出斐波那契数列的前20个数
	fmt.Println("Fibonacci Series:")
	fibonacci := []int{0, 1}
	for i := 2; i < 20; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// 4. 定义一个结构体表示学生（包含姓名和年龄），创建一个学生对象并打印
	type Student struct {
		Name string
		Age  int
	}
	student := Student{Name: "Alice", Age: 18}
	fmt.Println("Student:", student)

	// 5. 使用函数求两个整数的和并返回结果
	fmt.Println("Addition of 5 and 3:", add(5, 3))
}

func add(a, b int) int {
	return a + b
}