package main

import "fmt"

func main() {
    // 1. 反转字符串
    str := "Hello, World!"
    reversedStr := ""
    for i := len(str)-1; i >= 0; i-- {
      reversedStr += string(str[i])
    }
    fmt.Println(reversedStr)

    // 2. 判断一个数是否为素数
    num := 29
    isPrime := true
    for i := 2; i <= num/2; i++ {
      if num%i == 0 {
        isPrime = false
        break
      }
    }
    fmt.Println(isPrime)

    // 3. 求最大公约数和最小公倍数
    a, b := 54, 72
    gcd := 1
    lcm := a * b
    for i := 1; i <= a && i <= b; i++ {
      if a%i == 0 && b%i == 0 {
        gcd = i
      }
    }
    lcm /= gcd
    fmt.Println(gcd, lcm)

    // 4. 计算斐波那契数列
    n := 10
    fib := make([]int, n)
    fib[0] = 0
    fib[1] = 1
    for i := 2; i < n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    fmt.Println(fib)
}
```