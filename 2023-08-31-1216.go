package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个数的和
	var a, b int = 3, 4
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 5
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：找出数组中的最大值
	arr := []int{9, 3, 6, 1, 7, 2, 5, 8, 4}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("Max:", max)

	// 练习5：计算斐波那契数列
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci:", fibonacci)
}