package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算1到10的和
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("总和：", sum)

    // 练习4：判断一个数是不是质数
    number := 17
    isPrime := true
    for i := 2; i <= number/2; i++ {
        if number%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(number, "是质数")
    } else {
        fmt.Println(number, "不是质数")
    }

    // 练习5：判断一个数组中的元素是否都是偶数
    nums := []int{2, 4, 6, 8, 10}
    allEven := true
    for _, num := range nums {
        if num%2 != 0 {
            allEven = false
            break
        }
    }
    if allEven {
        fmt.Println("数组中的元素都是偶数")
    } else {
        fmt.Println("数组中的元素不全是偶数")
    }
}