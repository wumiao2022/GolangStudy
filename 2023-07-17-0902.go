package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：交换两个变量的值
	a := 5
	b := 10
	// 方法1：使用中间变量
	temp := a
	a = b
	b = temp
	fmt.Printf("After swapping, a is %d and b is %d\n", a, b)
	// 方法2：使用加减法
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After swapping again, a is %d and b is %d\n", a, b)

	// 练习4：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// 练习5：打印1到100之间所有的奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}
}