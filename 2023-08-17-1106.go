package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量和常量
	var name string = "Bob"
	age := 25
	const pi float64 = 3.1415926

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Pi:", pi)

	// 练习3：数组和切片
	grades := [5]int{98, 76, 85, 90, 88}
	fmt.Println("Grades:", grades)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// 练习4：循环和条件语句
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}

	// 练习5：函数和返回值
	sum := add(3, 5)
	fmt.Println("Sum:", sum)
}

func add(a, b int) int {
	return a + b
}