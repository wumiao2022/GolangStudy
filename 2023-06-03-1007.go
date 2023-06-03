package main

import "fmt"

func main() {
	// 练习1：打印“Hello, World!”
	fmt.Println("Hello, World!")

	// 练习2：计算两数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("sum is:", sum)

	// 练习3：判断奇偶性
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// 练习4：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习5：将切片中所有元素都乘以2
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
	fmt.Println(nums)
}