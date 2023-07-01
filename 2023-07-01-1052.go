package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算两个数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否是偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：计算斐波那契数列的前n项
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci:", fibonacci)

	// 练习6：求一个数组的最大值和最小值
	nums := []int{9, 2, 5, 1, 7, 4, 6, 8, 3}
	max := nums[0]
	min := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}