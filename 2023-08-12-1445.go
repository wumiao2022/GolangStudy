package main

import "fmt"

func main() {
	// 练习1：打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印出0到100之间所有的偶数
	for j := 0; j <= 100; j += 2 {
		fmt.Println(j)
	}

	// 练习3：计算1到100之间所有整数的和
	sum := 0
	for k := 1; k <= 100; k++ {
		sum += k
	}
	fmt.Println(sum)

	// 练习4：使用for循环实现一个九九乘法表
	for m := 1; m <= 9; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d*%d=%d ", n, m, n*m)
		}
		fmt.Println()
	}
}