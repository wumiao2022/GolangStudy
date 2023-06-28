package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// 练习3：判断一个数是否为偶数
	num := 17
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}

	// 练习4：计算一个数组中的元素和
	numbers := []int{1, 2, 3, 4, 5}
	sum = 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("The sum of the numbers is", sum)
}