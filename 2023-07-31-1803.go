package main

import "fmt"

// 练习1：编写一个函数，接收一个整数切片并返回切片中的最大值和最小值
func findMinMax(nums []int) (int, int) {
    min, max := nums[0], nums[0]
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

// 练习2：编写一个函数，接收一个整数切片并返回切片中的所有奇数
func findOdd(nums []int) []int {
    odds := []int{}
    for _, num := range nums {
        if num%2 != 0 {
            odds = append(odds, num)
        }
    }
    return odds
}

// 练习3：编写一个函数，接收一个字符串切片，并返回切片中所有元素的长度之和
func calculateTotalLength(strs []string) int {
    totalLength := 0
    for _, str := range strs {
        totalLength += len(str)
    }
    return totalLength
}

func main() {
    // 练习1的测试
    nums := []int{5, 3, 9, 1, 2, 8, 4, 7, 6}
    min, max := findMinMax(nums)
    fmt.Println("Min:", min, "Max:", max)

    // 练习2的测试
    nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    odds := findOdd(nums)
    fmt.Println("Odds:", odds)

    // 练习3的测试
    strs := []string{"hello", "world", "golang"}
    totalLength := calculateTotalLength(strs)
    fmt.Println("Total Length:", totalLength)
}
