package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. 将一个字符串逆序输出
	str := "Hello, World!"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	// 2. 计算一个数的平方根
	num := 16.0
	sqrt := math.Sqrt(num)
	fmt.Printf("The square root of %.2f is %.2f\n", num, sqrt)

	// 3. 求一个整数数组的和
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("The sum is", sum)
}