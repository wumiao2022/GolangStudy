package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：交换变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)

	// 练习3：求出1-100之间的所有奇数和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Printf("1-100之间的所有奇数和为%d\n", sum)

	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是素数\n", num)
	} else {
		fmt.Printf("%d不是素数\n", num)
	}

	// 练习5：打印杨辉三角第10行
	n := 10
	a := make([]int, n)
	a[0] = 1
	for i := 1; i < n; i++ {
		a[i] = 1
		for j := i - 1; j > 0; j-- {
			a[j] = a[j] + a[j-1]
		}
	}
	fmt.Println(a)
}