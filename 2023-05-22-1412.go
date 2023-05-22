package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")
    
    // 2. 输出1到100的所有奇数
    for i := 1; i <= 100; i++ {
        if i%2 == 1 {
            fmt.Println(i)
        }
    }
    
    // 3. 计算两个数的和并输出结果
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("The sum of", num1, "and", num2, "is", sum)
    
    // 4. 将一段字符串反转输出
    str := "Hello, World!"
    reversedStr := ""
    for i := len(str)-1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    fmt.Println(reversedStr)
    
    // 5. 实现一个函数，输入一个整数n，输出1到n的所有数字的和
    n := 10
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    fmt.Println("The sum of 1 to", n, "is", sum)
}