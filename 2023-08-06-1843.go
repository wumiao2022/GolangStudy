package main

import "fmt"

func main() {
	// 练习1: 打印数字1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习2: 检查一个数字是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3: 找出一个切片中的最大值
	nums := []int{12, 34, 92, 18, 56}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)

	// 练习4: 交换两个变量的值
	a, b := 5, 8
	a, b = b, a
	fmt.Println(a, b)

	// 练习5: 将字符串反转
	str := "Hello, world!"
	reverseStr := []rune(str)
	for i, j := 0, len(reverseStr)-1; i < j; i, j = i+1, j-1 {
		reverseStr[i], reverseStr[j] = reverseStr[j], reverseStr[i]
	}
	fmt.Println(string(reverseStr))
}