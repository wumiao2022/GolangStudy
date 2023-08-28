package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：交换两个变量的值
	a := 2
	b := 3
	a, b = b, a
	fmt.Printf("After swapping, a is %d and b is %d\n", a, b)

	// 练习4：打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：判断一个数是否是偶数
	num := 12
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}
}