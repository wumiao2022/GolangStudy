package main

import "fmt"

func main() {
	// 练习1: 计算两个整数的和
	a, b := 10, 20
	c := a + b
	fmt.Println(c)

	// 练习2: 判断一个数是奇数还是偶数
	num := 15
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习4: 交换两个整数的值
	x, y := 10, 20
	x, y = y, x
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	// 练习5: 求一个整数切片的和
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum =", sum)
}