package main

import "fmt"

func main() {
    // 实现一个函数，将两个整数相加，并返回结果
    func add(x, y int) int {
        return x + y
    }

    // 实现一个函数，判断一个整数是否为偶数，是则返回true，否则返回false
    func isEven(n int) bool {
        return n%2 == 0
    }

    // 实现一个函数，计算一个整数数组中所有元素的和，并返回结果
    func sum(nums []int) int {
        total := 0
        for _, num := range nums {
            total += num
        }
        return total
    }

    // 实现一个函数，返回字符串数组中长度最长的字符串
    func longestString(strs []string) string {
        maxLen := 0
        longestStr := ""
        for _, str := range strs {
            if len(str) > maxLen {
                maxLen = len(str)
                longestStr = str
            }
        }
        return longestStr
    }

    fmt.Println(add(3, 5))
    fmt.Println(isEven(7))
    fmt.Println(sum([]int{1, 2, 3, 4, 5}))
    fmt.Println(longestString([]string{"apple", "banana", "cherry"}))
}