package main

import "fmt"

func main() {
	// 练习1：输出1-10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：输出10-1的数字
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 练习3：输出1-100中的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：输出1-100中的奇数
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习5：求1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习6：求1-100中的偶数和
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习7：求1-100中的奇数和
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习8：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}