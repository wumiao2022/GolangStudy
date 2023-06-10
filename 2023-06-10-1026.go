package main

import (
	"fmt"
)

func main() {
	// 练习1：输出1到100的数字，但是如果是3的倍数，输出Fizz；如果是5的倍数，输出Buzz；如果同时是3和5的倍数，输出FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 练习2：将一个字符串反转，比如"hello world" -> "dlrow olleh"
	str := "hello world"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 练习3：将一个Int切片按照从小到大排序，比如[3, 2, 1, 5, 4] -> [1, 2, 3, 4, 5]
	nums := []int{3, 2, 1, 5, 4}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}