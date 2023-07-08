package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：判断一个整数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习4：求一个整数数组的和
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("和:", sum)

	// 练习5：计算斐波那契数列前20个数字
	fib := []int{0, 1}
	for i := 2; i < 20; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println("斐波那契数列前20个数字:", fib)
}