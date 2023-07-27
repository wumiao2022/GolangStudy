package main

import "fmt"

func main() {
	// 1. 请编写一个函数，接收两个整数参数，并将它们相加并返回结果

	func add(a, b int) int {
		return a + b
	}

	// 2. 请编写一个函数，接收一个整数数组作为参数，并计算并返回数组中所有元素的和

	func sum(nums []int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	// 3. 请编写一个函数，接收一个字符串切片作为参数，并将切片中的元素按照字典序排序并返回结果

	func sortStrings(strings []string) []string {
		sort.Strings(strings)
		return strings
	}

	// 4. 请编写一个函数，接收一个整数切片作为参数，并返回切片中所有元素的平均值

	func average(nums []int) float64 {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return float64(sum) / float64(len(nums))
	}

	// 5. 请编写一个函数，接收一个整数切片作为参数，并返回其中的最大值和最小值

	func findMinMax(nums []int) (int, int) {
		min := nums[0]
		max := nums[0]
		for _, num := range nums {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		return min, max
	}

	// 6. 请编写一个函数，接收一个字符串作为参数，并返回字符串的长度

	func getStringLength(str string) int {
		return len(str)
	}

	// 7. 请编写一个函数，接收一个整数作为参数，并判断该整数是否为偶数

	func isEven(num int) bool {
		return num%2 == 0
	}

	// 8. 请编写一个函数，接收一个整数作为参数，并判断该整数是否为质数

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

	// 9. 请编写一个函数，接收一个字符串作为参数，并判断该字符串是否为回文字符串

	func isPalindrome(str string) bool {
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-1-i] {
				return false
			}
		}
		return true
	}

	// 10. 请编写一个函数，接收一个整数切片作为参数，并返回一个新的切片，其中只包含奇数值

	func filterOdd(nums []int) []int {
		oddNums := []int{}
		for _, num := range nums {
			if num%2 != 0 {
				oddNums = append(oddNums, num)
			}
		}
		return oddNums
	}

	fmt.Println(add(2, 3))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(sortStrings([]string{"banana", "apple", "cherry"}))
	fmt.Println(average([]int{1, 2, 3, 4, 5}))
	fmt.Println(findMinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(getStringLength("Hello, world!"))
	fmt.Println(isEven(4))
	fmt.Println(isPrime(7))
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(filterOdd([]int{1, 2, 3, 4, 5}))
}