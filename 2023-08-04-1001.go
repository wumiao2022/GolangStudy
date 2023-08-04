package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习3：判断一个数是奇数还是偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习4：计算阶乘
	n := 5
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println("阶乘：", fact)
}
修改后的实例： 练习5：编写一个函数，将一个字符串逆序输出