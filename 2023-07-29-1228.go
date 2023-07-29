package main

import "fmt"

func main() {
	// 练习1：打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印出2的倍数，从2到100
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习3：打印出前20个斐波那契数列
	fibonacci := []int{0, 1}
	for i := 2; i <= 20; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// 练习4：计算1到100的所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}