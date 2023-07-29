package main

import (
	"fmt"
)

func main() {
	// 练习题1：求1到100之间所有奇数的和并打印
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("奇数的和为：", sum)

	// 练习题2：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习题3：判断一个数是否为素数并打印
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}
}

注意：这是一个完整的Go语言代码，你可以直接运行它来验证结果。每个练习在代码注释中有相应的说明。