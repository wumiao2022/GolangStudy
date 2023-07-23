package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("The sum of 1 to 100 is:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 17
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：打印一个九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}