package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量定义与使用
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 练习3：基本运算符
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// 练习4：for循环
	fmt.Println("Even numbers between 1 and 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习5：if-else条件语句
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习6：switch语句
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good!")
	case "C":
		fmt.Println("Average!")
	default:
		fmt.Println("Fail!")
	}

	// 练习7：函数定义与调用
	result := add(3, 5)
	fmt.Println("Sum:", result)

	// 练习8：数组和切片
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	slice := numbers[1:4]
	fmt.Println("Slice:", slice)

	// 练习9：指针
	x := 10
	y := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", y)
	fmt.Println("Value stored at address:", *y)
}

func add(a, b int) int {
	return a + b
}
