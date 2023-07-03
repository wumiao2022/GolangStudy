package main

import "fmt"

func main() {
    // 练习1: 输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3: 给定一个切片，反转切片中的元素
    slice := []int{1, 2, 3, 4, 5}
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
    fmt.Println(slice)

    // 练习4: 计算斐波那契数列的前n个数
    n := 10
    fibonacci := make([]int, n)
    fibonacci[0], fibonacci[1] = 0, 1
    for i := 2; i < n; i++ {
        fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
    }
    fmt.Println(fibonacci)

    // 练习5: 检查一个数字是否为素数
    num := 7
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(isPrime)
}
