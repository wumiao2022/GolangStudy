package main

import "fmt"

// 练习1：打印数字1到10
func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}

// 练习2：计算1到100之间所有偶数的和
func main() {
    sum := 0
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            sum += i
        }
    }
    fmt.Println(sum)
}

// 练习3：判断一个字符串是否为回文字符串
func main() {
    str := "level"
    isPalindrome := true
    for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
        if str[i] != str[j] {
            isPalindrome = false
            break
        }
    }
    fmt.Println(isPalindrome)
}

// 练习4：打印九九乘法表
func main() {
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d\t", i, j, i*j)
        }
        fmt.Println()
    }
}

// 练习5：计算斐波那契数列的前20项
func main() {
    fibonacci := []int{0, 1}
    for i := 2; i < 20; i++ {
        fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
    }
    fmt.Println(fibonacci)
}