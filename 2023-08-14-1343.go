package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")
    
    // 练习2：计算10的阶乘
    factorial := 1
    for i := 1; i <= 10; i++ {
        factorial *= i
    }
    fmt.Println("10的阶乘为：", factorial)
    
    // 练习3：判断一个数是否为质数
    number := 17
    prime := true
    for i := 2; i <= number/2; i++ {
        if number%i == 0 {
            prime = false
            break
        }
    }
    if prime {
        fmt.Println(number, "是质数")
    } else {
        fmt.Println(number, "不是质数")
    }
    
    // 练习4：计算斐波那契数列的前20项
    fibonacci := []int{0, 1}
    for i := 2; i < 20; i++ {
        fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
    }
    fmt.Println("斐波那契数列的前20项为：", fibonacci)
}