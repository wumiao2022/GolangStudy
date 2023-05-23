package main

import "fmt"

func main() {
	// 声明变量
	var num int

	// 获取用户输入
	fmt.Println("请输入一个正整数:")
	fmt.Scan(&num)

	// 判断是否是素数
	if isPrime(num) {
		fmt.Printf("%d是素数\n", num)
	} else {
		fmt.Printf("%d不是素数\n", num)
	}
}

// 判断是否是素数的函数
func isPrime(num int) bool {
	if num <= 1 {
		return false
	} else if num == 2 {
		return true
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
}