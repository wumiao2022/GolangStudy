package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

	// 练习4：循环打印1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个切片并遍历
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("The number at index %d is %d\n", i, num)
	}
}