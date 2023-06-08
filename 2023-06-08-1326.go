package main

import "fmt"

func main() {
	//练习1：输出1-10所有的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//练习2：判断一个年份是否是闰年
	year := 2020
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(year, "年是闰年")
	} else {
		fmt.Println(year, "年不是闰年")
	}

	//练习3：计算斐波那契数列的第n项（n>=3）
	n := 10
	fib1, fib2 := 1, 1
	for i := 3; i <= n; i++ {
		fib1, fib2 = fib2, fib1+fib2
	}
	fmt.Println("斐波那契数列的第", n, "项是：", fib2)

	//练习4：求1-100完数的个数和每个完数
	fmt.Println("1-100的完数有：")
	for i := 1; i <= 100; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			fmt.Println(i)
		}
	}
}