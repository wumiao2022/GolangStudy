package main

import "fmt"

func main() {
	// 1. 将字符串反转
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 2. 打印乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
		}
		fmt.Println()
	}

	// 3. 检查字符串是否为回文
	str2 := "level"
	isPalindrome := true
	for i := 0; i < len(str2)/2; i++ {
		if str2[i] != str2[len(str2)-1-i] {
			isPalindrome = false
			break
		}
	}
	fmt.Println(isPalindrome)

	// 4. 将切片中的元素按照从大到小的顺序排列
	nums := []int{5, 2, 7, 1, 9}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
}