package main

import "fmt"

func main() {
	// 1. 打印1-100的所有奇数
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 2. 计算1到100的和，当和大于等于100时输出结果并终止程序
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 100 {
			fmt.Println("Sum is:", sum)
			break
		}
	}

	// 3. 打印出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 4. 判断101-200之间有多少个素数
	// 素数：除了1和本身，不能被其他自然数整除的数
	count := 0
	for i := 101; i <= 200; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}
	fmt.Println("There are", count, "prime numbers between 101 and 200.")

	// 5. 找出1-1000中的水仙花数
	// 水仙花数：一个三位数，其各位数字立方和等于该数本身
	for i := 100; i < 1000; i++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}