package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：变量赋值与打印
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Printf("num1 = %d\n", num1)
	fmt.Printf("num2 = %.2f\n", num2)
	fmt.Printf("str = %s\n", str)

	// 练习3：条件语句
	num := 5
	if num > 0 {
		fmt.Println("Number is positive")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// 练习4：循环语句
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习5：数组与遍历
	numbers := [3]int{1, 2, 3}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 练习6：切片与遍历
	numbersSlice := []int{1, 2, 3, 4, 5}
	for _, value := range numbersSlice {
		fmt.Println(value)
	}

	// 练习7：函数定义与调用
	result := sum(10, 20)
	fmt.Println("Sum is", result)
}

// 函数：求和
func sum(a, b int) int {
	return a + b
}