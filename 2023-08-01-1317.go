package main

import "fmt"

func main() {
	// 实现一个函数，判断一个整数是否是奇数
	num := 5
	isOdd := func(n int) bool {
		if n%2 != 0 {
			return true
		}
		return false
	}(num)

	fmt.Printf("Is %d odd? %t\n", num, isOdd)
}