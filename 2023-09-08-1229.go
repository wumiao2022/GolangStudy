package main

import "fmt"

func main() {
    // 练习1：打印出1到10之间的所有偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习2：计算数组中所有元素的和
    nums := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("Sum:", sum)

    // 练习3：计算1到100之间所有能被3整除的数字的平均值
    count := 0
    total := 0
    for i := 1; i <= 100; i++ {
        if i%3 == 0 {
            total += i
            count++
        }
    }
    average := float64(total) / float64(count)
    fmt.Println("Average:", average)

    // 练习4：判断一个字符串是否是回文字符串
    word := "level"
    isPalindrome := true
    for i := 0; i < len(word)/2; i++ {
        if word[i] != word[len(word)-1-i] {
            isPalindrome = false
            break
        }
    }
    fmt.Println("Is palindrome?", isPalindrome)
}