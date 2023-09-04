package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算并打印1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否为偶数
	num := 4
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：计算并打印Fibonacci数列的前n个数字
	n := 10
	prev := 0
	next := 1
	fmt.Println(prev) // 打印第一个数字0
	fmt.Println(next) // 打印第二个数字1
	for i := 2; i < n; i++ {
		// 计算下一个数字，并更新prev和next的值
		sum := prev + next
		fmt.Println(sum)
		prev = next
		next = sum
	}
}