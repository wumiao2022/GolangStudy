package main

import "fmt"

func main() {
    //1. 输出10个"Hello, World!"
    for i := 1; i <= 10; i++ {
        fmt.Println("Hello, World!")
    }

    //2. 求1-100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1-100的和为：", sum)

    //3. 遍历一个切片并输出其中的每个元素
    slice := []string{"apple", "banana", "orange"}
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    //4. 判断一个数是否是素数
    num := 17
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

    //5. 实现一个函数，计算两个整数的和
    fmt.Println("1 + 2 = ", add(1, 2))
}

func add(a, b int) int {
    return a + b
}