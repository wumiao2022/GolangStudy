package main

import (
	"fmt"
)

func main() {
	// 练习实例1：打印数组元素
	arr1 := []int{1, 2, 3, 4, 5}
	for _, num := range arr1 {
		fmt.Println(num)
	}

	// 练习实例2：求两个数的和
	sum := 0
	num1 := 10
	num2 := 20
	sum = num1 + num2
	fmt.Println(sum)

	// 练习实例3：判断某个数是否是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习实例4：打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}