package main

import "fmt"

func main() {
    // 练习1：打印出"Hello, World!"
    fmt.Println("Hello, World!")
    
    // 练习2：打印出整数1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 练习3：打印出偶数1到20
    for i := 1; i <= 20; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
    
    // 练习4：计算1到100的和并输出结果
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)
    
    // 练习5：判断一个数是否是素数（只能被1和自身整除）
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
        if num % i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Printf("%d is prime: %v\n", num, isPrime)
}