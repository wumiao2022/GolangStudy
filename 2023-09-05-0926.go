package main

import "fmt"

// 练习1: 打印Hello, World!
func main() {
	fmt.Println("Hello, World!")
}

// 练习2: 定义一个函数，将输入的整数乘以2后返回
func double(num int) int {
	return num * 2
}

// 练习3: 定义一个函数，将两个整数相加后返回结果
func add(num1, num2 int) int {
	return num1 + num2
}

// 练习4: 定义一个函数，判断一个整数是否为偶数，返回布尔值
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// 练习5: 定义一个函数，接收一个整数数组，将其中的奇数元素返回一个新的数组
func filterOddNumbers(nums []int) []int {
	var result []int
	for _, num := range nums {
		if num%2 != 0 {
			result = append(result, num)
		}
	}
	return result
}
// ... 还有更多的练习示例代码，更多练习请继续查看其他代码
```