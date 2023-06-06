package main

import (
	"fmt"
)

func main() {
	// 练习一：打印数字1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习二：打印乘法口诀表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习三：计算1~100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1~100的和为：", sum)

	// 练习四：计算1~100的偶数和
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("1~100的偶数和为：", sum)

	// 练习五：计算斐波那契数列的第10项
	f1, f2 := 1, 1
	for i := 3; i <= 10; i++ {
		f1, f2 = f2, f1+f2
	}
	fmt.Println("第10项斐波那契数列的值为：", f2)
}