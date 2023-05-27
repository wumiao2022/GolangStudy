package main

import "fmt"

func main() {
	// 1. 打印乘法口诀表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}

	// 2. 求100以内的素数
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j <= i-1; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d ", i)
		}
	}

	// 3. 将一个正整数分解质因数
	num := 36
	for i := 2; i <= num; i++ {
		for num%i == 0 {
			fmt.Printf("%d ", i)
			num = num / i
		}
	}
	if num > 1 {
		fmt.Printf("%d ", num)
	}

	// 4. 逆序打印数字
	num = 123456789
	for num > 0 {
		fmt.Printf("%d", num%10)
		num = num / 10
	}
}