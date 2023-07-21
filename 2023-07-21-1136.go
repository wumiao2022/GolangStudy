package main

import "fmt"

func main() {
	// 1. 编写一个函数，将两个整数相加并返回结果
	func add(x, y int) int {
		return x + y
	}

	// 2. 编写一个函数，判断一个数是否为偶数，是则返回true，否则返回false
	func isEven(num int) bool {
		return num%2 == 0
	}

	// 3. 编写一个函数，计算一个整数数组中所有元素的和，并返回结果
	func sum(nums []int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	// 4. 编写一个函数，判断一个字符串是否为回文字符串，是则返回true，否则返回false
	func isPalindrome(str string) bool {
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-1-i] {
				return false
			}
		}
		return true
	}

	// 5. 编写一个函数，将一个字符串逆序并返回
	func reverseString(str string) string {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	// 6. 编写一个函数，计算一个整数数组中的最大值并返回该值
	func max(nums []int) int {
		max := nums[0]
		for _, num := range nums {
			if num > max {
				max = num
			}
		}
		return max
	}

	// 7. 编写一个函数，判断两个字符串是否为同构字符串，是则返回true，否则返回false
	func isIsomorphic(str1, str2 string) bool {
		if len(str1) != len(str2) {
			return false
		}

		mapping := make(map[byte]byte)
		usedChars := make(map[byte]bool)

		for i := 0; i < len(str1); i++ {
			char1 := str1[i]
			char2 := str2[i]

			if val, ok := mapping[char1]; ok {
				if val != char2 {
					return false
				}
			} else {
				if usedChars[char2] {
					return false
				}
				mapping[char1] = char2
				usedChars[char2] = true
			}
		}

		return true
	}

	// 调用以上函数进行测试
	fmt.Println(add(1, 2))
	fmt.Println(isEven(4))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(isPalindrome("level"))
	fmt.Println(reverseString("hello"))
	fmt.Println(max([]int{1, 2, 3, 4, 5}))
	fmt.Println(isIsomorphic("egg", "add"))
}