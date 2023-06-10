package main

import "fmt"

func main() {
    // 练习题1
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(arr))

    // 练习题2
    fmt.Println(fibonacci(10))

    // 练习题3
    strs := []string{"hello", "world", "golang", "go", "practice"}
    fmt.Println(longestString(strs))

    // 练习题4
    fmt.Println(reverseString("hello world"))

    // 练习题5
    fmt.Println(isPalindrome("racecar"))
}

// 练习题1：计算一个整数数组的和
func sum(arr []int) int {
    total := 0
    for _, num := range arr {
        total += num
    }
    return total
}

// 练习题2：生成一个斐波那契数列
func fibonacci(n int) []int {
    if n == 0 {
        return []int{}
    } else if n == 1 {
        return []int{0}
    } else if n == 2 {
        return []int{0, 1}
    } else {
        arr := make([]int, n)
        arr[0] = 0
        arr[1] = 1
        for i := 2; i < n; i++ {
            arr[i] = arr[i-1] + arr[i-2]
        }
        return arr
    }
}

// 练习题3：找出一个字符串数组中长度最长的字符串
func longestString(strs []string) string {
    longest := ""
    for _, str := range strs {
        if len(str) > len(longest) {
            longest = str
        }
    }
    return longest
}

// 练习题4：反转一个字符串
func reverseString(str string) string {
    result := ""
    for i := len(str) - 1; i >= 0; i-- {
        result += string(str[i])
    }
    return result
}

// 练习题5：判断一个字符串是否是回文字符串
func isPalindrome(str string) bool {
    for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
        if str[i] != str[j] {
            return false
        }
    }
    return true
}