package main

import "fmt"

func main() {
	// 1. 交换两个变量的值
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)

	// 2. 判断一个数是否是质数
	n := 7
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 3. 计算斐波那契数列的第n项
	n = 8
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	fmt.Println(a)
	
	// 4. 大整数相加
	num1 := "123456789123456789"
	num2 := "987654321987654321"
	res := ""
	carry := 0
	n1, n2 := len(num1), len(num2)
	for i, j := n1-1, n2-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		sum := a + b + carry
		res = fmt.Sprintf("%d%s", sum%10, res)
		carry = sum / 10 
	}
	if carry != 0 {
		res = fmt.Sprintf("%d%s", carry, res)
	}
	fmt.Println(res)
}