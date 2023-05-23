package main

import (
	"fmt"
)

func main() {
	// 练习1：输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算 1+2+...+10 的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1+2+...+10 =", sum)

	// 练习3：判断一个数是不是质数
	num := 23
	flag := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println(num, "是质数")
	} else {
		fmt.Println(num, "不是质数")
	}

	// 练习4：求斐波那契数列的第n项，n由用户输入
	var n int
	fmt.Print("请输入要求的斐波那契数列的第几项：")
	fmt.Scan(&n)
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	fmt.Printf("斐波那契数列的第%d项是%d\n", n, b)

	// 练习5：交换两个变量的值
	a, b = b, a
	fmt.Println("现在a的值是", a, "，b的值是", b)
}