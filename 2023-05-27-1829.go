package main

import (
	"fmt"
)

// 1. 求一个数的平方值
func square(num int) int {
	return num * num
}

// 2. 判断一个数是否是奇数
func isOdd(num int) bool {
	return num%2 != 0
}

// 3. 判断一个数是否是质数
func isPrime(num int) bool {
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

// 4. 按照升序返回两个数的最大公约数和最小公倍数
func gcdAndLcm(num1, num2 int) (int, int) {
	var gcd, lcm int
	var maxNum int
	if num1 > num2 {
		maxNum = num1
	} else {
		maxNum = num2
	}

	for i := 1; i <= maxNum; i++ {
		if num1%i == 0 && num2%i == 0 {
			gcd = i
		}
	}

	lcm = num1 * num2 / gcd
	return gcd, lcm
}

// 5. 将一个整数划分为若干个奇数之和的形式，如：7 = 1+3+3，13 = 1+3+3+3+3
func splitNum(num int) []int {
	res := []int{}
	for num > 0 {
		if isOdd(num) {
			res = append(res, 1)
			num--
		} else {
			res = append(res, 2)
			num -= 2
		}
	}
	return res
}

// 6. 实现冒泡排序算法
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 7. 实现插入排序算法
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

// 8. 实现选择排序算法
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// 9. 给定一个字符串，判断是否是回文字符串
func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// 10. 实现一个类似 Python range() 函数功能的函数
func myRange(start, stop, step int) []int {
	res := []int{}
	for i := start; i < stop; i += step {
		res = append(res, i)
	}
	return res
}

func main() {
	fmt.Println(square(5))
	fmt.Println(isOdd(6))
	fmt.Println(isPrime(7))
	fmt.Println(gcdAndLcm(18, 24))
	fmt.Println(splitNum(13))
	fmt.Println(bubbleSort([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}))
	fmt.Println(insertionSort([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}))
	fmt.Println(selectionSort([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}))
	fmt.Println(isPalindrome("abcdcba"))
	fmt.Println(myRange(0, 10, 3))
}