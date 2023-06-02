package main

import "fmt"

func main() {
    // 1. sum函数，计算两个数的和
    func sum(a, b int) int {
        return a + b
    }
    fmt.Println(sum(2, 3)) // 输出：5

    // 2. 判断是否是闰年
    func isLeapYear(year int) bool {
        if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
            return true
        } else {
            return false
        }
    }
    fmt.Println(isLeapYear(2020)) // 输出：true
    fmt.Println(isLeapYear(2021)) // 输出：false

    // 3. 反转一个字符串
    func reverse(s string) string {
        r := []rune(s)
        for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
            r[i], r[j] = r[j], r[i]
        }
        return string(r)
    }
    fmt.Println(reverse("hello")) // 输出：olleh

    // 4. 计算一个整数的阶乘
    func factorial(n int) int {
        if n == 0 {
            return 1
        }
        return n * factorial(n-1)
    }
    fmt.Println(factorial(5)) // 输出：120

    // 5. 判断一个数是不是质数
    func isPrime(n int) bool {
        if n <= 1 {
            return false
        }
        for i := 2; i*i <= n; i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }
    fmt.Println(isPrime(7)) // 输出：true
    fmt.Println(isPrime(9)) // 输出：false
}