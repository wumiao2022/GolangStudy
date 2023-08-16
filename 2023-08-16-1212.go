package main
 
import (
    "fmt"
)
 
func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")
 
    // 练习2：输出1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
 
    // 练习3：计算两个整数的和
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)
 
    // 练习4：判断一个数是否为偶数
    num := 6
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
 
    // 练习5：将字符串反转输出
    str := "Hello, World!"
    reversedStr := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    fmt.Println("Reversed string:", reversedStr)
}