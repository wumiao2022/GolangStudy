package main

import "fmt"

func main() {
    // 实现一个简单的计算器，支持加、减、乘、除四种运算
    num1 := 10.0
    num2 := 5.0

    fmt.Println("Addition:", num1+num2)
    fmt.Println("Subtraction:", num1-num2)
    fmt.Println("Multiplication:", num1*num2)
    fmt.Println("Division:", num1/num2)

    // 实现一个简单的交换变量的函数
    a := 1
    b := 2
    a, b = b, a
    fmt.Println("a:", a, " b:", b)

    // 实现一个函数，计算斐波那契数列的值
    result := fibonacci(10)
    fmt.Println("Fibonacci result:", result)

    // 实现一个函数，找到数组中的最大值
    arr := []int{1, 5, 3, 8, 2}
    max := getMax(arr)
    fmt.Println("Max value:", max)

    // 实现一个函数，判断一个字符串是否是回文字符串
    str := "level"
    isPalindrome := checkPalindrome(str)
    fmt.Println("Is palindrome:", isPalindrome)

}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func getMax(arr []int) int {
    max := arr[0]
    for _, value := range arr {
        if value > max {
            max = value
        }
    }
    return max
}

func checkPalindrome(str string) bool {
    length := len(str)
    for i := 0; i < length/2; i++ {
        if str[i] != str[length-1-i] {
            return false
        }
    }
    return true
}