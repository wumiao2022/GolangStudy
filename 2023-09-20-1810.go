package main

import "fmt"

func main() {
	// 练习1：打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习2：计算斐波那契数列前20个数
	fib := make([]int, 20)
	fib[0], fib[1] = 0, 1
	for i := 2; i < 20; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println("斐波那契数列前20个数:", fib)

	// 练习3：反转字符串
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("反转后的字符串:", reversedStr)

	// 练习4：找出数组中的最大值
	nums := []int{2, 9, 5, 7, 3, 1, 8, 6, 4}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println("数组中的最大值:", max)
}
