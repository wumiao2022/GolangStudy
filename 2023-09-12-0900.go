package main

import "fmt"

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")
    
    // 练习2：计算1到10的和
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("1到10的和为：", sum)
    
    // 练习3：判断一个数是否为素数
    var num int = 17
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "是素数")
    } else {
        fmt.Println(num, "不是素数")
    }
    
    // 练习4：输出斐波那契数列前20个数
    fibonacci := []int{0, 1}
    for i := 2; i < 20; i++ {
        fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
    }
    fmt.Println("斐波那契数列前20个数为：", fibonacci)
    
    // 练习5：判断一个字符串是否为回文字符串
    var str string = "level"
    isPalindrome := true
    for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-i-1] {
            isPalindrome = false
            break
        }
    }
    if isPalindrome {
        fmt.Println(str, "是回文字符串")
    } else {
        fmt.Println(str, "不是回文字符串")
    }
}