package main

import "fmt"

func main() {
	// 1. 编写一个函数，传入一个整数n，打印输出从1到n之间的所有偶数
	printEvenNumbers(10)

	// 2. 编写一个函数，传入一个字符串，将其按照相反的顺序输出
	reverseString("Hello, World!")

	// 3. 编写一个函数，传入一个整数n，计算并返回n的阶乘
	fmt.Println(factorial(5))

	// 4. 编写一个函数，传入一个数组arr和一个目标值target，在arr中查找两个元素的和等于target，返回这两个元素的下标
	arr := []int{2, 7, 11, 15}
	target := 9
	indices := twoSum(arr, target)
	fmt.Println(indices)
}

func printEvenNumbers(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func reverseString(s string) {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func twoSum(arr []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range arr {
		complement := target - num
		if index, found := numsMap[complement]; found {
			return []int{index, i}
		}
		numsMap[num] = i
	}
	return []int{}
}