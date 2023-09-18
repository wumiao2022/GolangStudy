package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 变量赋值与打印
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)
	fmt.Println("str =", str)

	// 练习3: 算术运算
	num3 := 5
	num4 := 2
	sum := num3 + num4
	difference := num3 - num4
	product := num3 * num4
	quotient := float64(num3) / float64(num4)
	remainder := num3 % num4
	fmt.Println("sum =", sum)
	fmt.Println("difference =", difference)
	fmt.Println("product =", product)
	fmt.Println("quotient =", quotient)
	fmt.Println("remainder =", remainder)

	// 练习4: 条件判断
	if num1 > num3 {
		fmt.Println("num1 is greater than num3")
	} else {
		fmt.Println("num3 is greater than num1")
	}

	// 练习5: 循环语句
	for i := 1; i <= 5; i++ {
		fmt.Println("Current number:", i)
	}

	// 练习6: 数组与切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	// 练习7: 函数
	result := add(3, 4)
	fmt.Println("Sum:", result)
}

func add(a, b int) int {
	return a + b
}