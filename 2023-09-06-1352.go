package main

import (
	"fmt"
)

func main() {
	// 1. 将一个字符串逆序输出
	str := "Hello, World!"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}

	// 2. 求一个整数数组的和
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("\nSum:", sum)

	// 3. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 4. 判断一个年份是否为闰年
	year := 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}

	// 5. 求斐波那契数列的前n项
	n := 10
	fibonacci := make([]int, n)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println("Fibonacci Sequence:", fibonacci)
}