package main

import "fmt"

func main() {
	// 练习1: 定义一个函数，实现两个整数相加的功能，并返回相加的结果
	func addTwoNumbers(a, b int) int {
		return a + b
	}

	// 练习2: 定义一个函数，判断一个整数是否为偶数，是则返回true，否则返回false
	func isEvenNumber(num int) bool {
		if num%2 == 0 {
			return true
		} else {
			return false
		}
	}

	// 练习3: 定义一个函数，接收一个字符串作为参数，返回字符串的长度
	func getStringLength(str string) int {
		return len(str)
	}

	fmt.Println(addTwoNumbers(4, 6))        // 输出: 10
	fmt.Println(isEvenNumber(7))            // 输出: false
	fmt.Println(getStringLength("Golang"))  // 输出: 6
}