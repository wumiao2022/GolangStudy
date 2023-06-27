package main

import "fmt"

func main() {
	// 实例1：输出Hello World
	fmt.Println("Hello World")

	// 实例2：输出1~10之间的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 实例3：求1~100之间所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 实例4：输出一张4行5列，元素为0~19的二维数组
	arr := [4][5]int{}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = i*5 + j
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}