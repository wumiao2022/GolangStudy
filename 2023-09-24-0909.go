package main

import "fmt"

func main() {
	// 练习：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习：将两个数相加输出结果
	a := 3
	b := 4
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习：将一个字符串反转输出
	str := "Hello"
	reverseStr := reverseString(str)
	fmt.Println("Reverse:", reverseStr)

	// 练习：给定一个数组，找出数组中的最大值和最小值
	arr := []int{5, 2, 8, 1, 9, 3}
	max, min := findMaxAndMin(arr)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)

	// 练习：判断一个数是否为素数
	num := 17
	isPrime := isPrimeNumber(num)
	fmt.Println("Is Prime:", isPrime)
}

// 反转字符串函数
func reverseString(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

// 查找数组中的最大值和最小值函数
func findMaxAndMin(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}

// 判断素数函数
func isPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}