package main

import "fmt"

func main() {
    // 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

    // 练习2: 计算两个整数的和并输出结果
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)
    
    // 练习3: 定义一个字符串变量并输出
    message := "Hello from Golang!"
    fmt.Println(message)
    
    // 练习4: 使用for循环输出1至10的平方值
    for i := 1; i <= 10; i++ {
        square := i * i
        fmt.Println("Square of", i, ":", square)
    }
    
    // 练习5: 判断一个数字是否为偶数，并输出结果
    num := 15
    if num % 2 == 0 {
        fmt.Println(num, "is even.")
    } else {
        fmt.Println(num, "is odd.")
    }
    
    // 练习6: 定义一个切片并输出其中的元素
    numbers := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(numbers); i++ {
        fmt.Println("Element", i, ":", numbers[i])
    }
}