package main

func main() {
    // 1. 打印Hello World
    println("Hello World")
    
    // 2. 计算两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    println("Sum:", sum)
    
    // 3. 判断一个数是否为偶数
    num := 25
    if num%2 == 0 {
        println(num, "is even")
    } else {
        println(num, "is odd")
    }
    
    // 4. 反转字符串
    str := "Hello, Golang"
    reversedStr := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    println("Reversed string:", reversedStr)
}