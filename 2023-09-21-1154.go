package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：两数相加
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断数字是奇数还是偶数
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 练习4：计算平均值
	numbers := []float64{1.2, 2.4, 3.6, 4.8}
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	average := total / float64(len(numbers))
	fmt.Println("Average:", average)

	// 练习5：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}