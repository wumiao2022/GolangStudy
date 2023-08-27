package main

func main() {
	// 练习1：打印Hello, World!
	println("Hello, World!")

	// 练习2：使用循环打印10次"Hello"
	for i := 0; i < 10; i++ {
		println("Hello")
	}

	// 练习3：计算两个整数相加的结果并打印
	num1 := 5
	num2 := 3
	sum := num1 + num2
	println("Sum:", sum)

	// 练习4：使用数组存储10个整数并打印出来
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		println(num)
	}

	// 练习5：使用切片存储10个整数并打印出来
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range slice {
		println(num)
	}
}