package main

import "fmt"

func main() {
	// 1. 打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 3. 判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 4. 找出一个数组中的最大值
	nums := []int{4, 9, 2, 7, 5}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println(max)

	// 5. 判断一个字符串是否是回文字符串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("是回文字符串")
	} else {
		fmt.Println("不是回文字符串")
	}
}