package main

import (
	"fmt"
)

func main() {
	//练习1：输出1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//练习2：输出10到1的整数
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	//练习3：求1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	//练习4：求1到100之间所有奇数的和
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	//练习5：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}