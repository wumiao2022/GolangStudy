package main

import "fmt"

func main() {
	// 1. 交换两个变量的值，并打印结果

	a := 10
	b := 20

	fmt.Printf("交换前，a的值为：%d，b的值为：%d\n", a, b)

	a, b = b, a

	fmt.Printf("交换后，a的值为：%d，b的值为：%d\n\n", a, b)

	// 2. 判断一个数是奇数还是偶数

	num := 17

	if num%2 == 0 {
		fmt.Printf("%d是偶数\n\n", num)
	} else {
		fmt.Printf("%d是奇数\n\n", num)
	}

	// 3. 判断一个年份是否为闰年

	year := 2024

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Printf("%d年是闰年\n\n", year)
	} else {
		fmt.Printf("%d年不是闰年\n\n", year)
	}

	// 4. 打印九九乘法表

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

}