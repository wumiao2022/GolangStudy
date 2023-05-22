package main

import "fmt"

// 题目一：两数之和
func twoSum(nums []int, target int) []int {
    var index []int
    hash := map[int]int{}
    for i, num := range nums {
        if val, ok := hash[target-num]; ok {
            index = []int{val, i}
            break
        }
        hash[num] = i
    }
    return index
}

// 题目二：回文数
func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    y := 0
    for x > y {
        y = y*10 + x%10
        x /= 10
    }
    return x == y || x == y/10
}

// 题目三：最长回文子串
func longestPalindrome(s string) string {
    runes := []rune(s)
    start, maxLen := 0, 0
    for i := 0; i < len(runes); {
        if len(runes)-i <= maxLen/2 {
            break
        }
        j, k := i, i
        for k < len(runes)-1 && runes[k+1] == runes[k] {
            k++
        }
        i = k + 1
        for k < len(runes)-1 && j > 0 && runes[k+1] == runes[j-1] {
            k++
            j--
        }
        newLen := k - j + 1
        if newLen > maxLen {
            start = j
            maxLen = newLen
        }
    }
    return string(runes[start : start+maxLen])
}

func main() {
    // 验证题目一：两数之和
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums, target)) // output: [0 1]

    // 验证题目二：回文数
    fmt.Println(isPalindrome(121)) // output: true
    fmt.Println(isPalindrome(-121)) // output: false

    // 验证题目三：最长回文子串
    fmt.Println(longestPalindrome("babad")) // output: bab
}