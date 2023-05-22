package main

import "fmt"

func main() {
	// 示例1：打印1~100的数字
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// 示例2：计算1+2+3+...+100的值
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 示例3：打印字符串"Hello, world!"10次
	str := "Hello, world!"
	for i := 1; i <= 10; i++ {
		fmt.Println(str)
	}

	// 示例4：打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 示例5：查找一个数是否在数组中，并返回下标
	nums := []int{10, 20, 30, 40, 50}
	target := 30
	for i, num := range nums {
		if num == target {
			fmt.Printf("下标为%d\n", i)
			break
		}
	}
}