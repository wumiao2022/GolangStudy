package main

import "fmt"

func main() {
    // 1. 打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 2. 打印0到50之间的偶数
    for i := 0; i <= 50; i += 2 {
        fmt.Println(i)
    }

    // 3. 计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 4. 将一个字符串反转
    str := "Hello, World!"
    reversedStr := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    fmt.Println(reversedStr)

    // 5. 判断一个数是否为素数
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(num, "is prime:", isPrime)
}