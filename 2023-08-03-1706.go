package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并打印
	sum := 1 + 2
	fmt.Println("1 + 2 =", sum)

	// 练习3：打印1到10的所有奇数
	fmt.Println("Odd numbers from 1 to 10:")
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习4：使用for循环打印九九乘法表
	fmt.Println("Multiplication table:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习5：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number.")
	} else {
		fmt.Println(num, "is not a prime number.")
	}
}

// 练习结束，以上为示例代码，可以运行并测试结果。