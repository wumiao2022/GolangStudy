package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：求1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否为素数
    num := 17
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(num, "is prime:", isPrime)

    // 练习4：交换两个变量的值
    a, b := 5, 10
    a, b = b, a
    fmt.Println("a:", a, "b:", b)

    // 练习5：冒泡排序
    arr := []int{8, 3, 1, 6, 2, 7, 4, 5}
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println("Sorted array:", arr)
}
