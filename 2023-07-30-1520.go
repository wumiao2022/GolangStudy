package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算两个整数的和并输出
	a := 5
	b := 7
	sum := a + b
	fmt.Println(sum)

	// 练习4：判断一个数是否为偶数并输出结果
	num := 8
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 练习5：输出一个九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
