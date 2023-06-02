package main

import "fmt"

func main() {
	// 练习1：打印1~10的正整数
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 练习2：打印1~10中的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// 练习3：打印1~10中的奇数
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// 练习4：求1~10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习5：求1~10中的偶数之和
	sum = 0
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习6：求1~10中的奇数之和
	sum = 0
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习7：打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习8：求1~100内的素数
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}