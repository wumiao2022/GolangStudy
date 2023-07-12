package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和
    a := 5
    b := 7
    sum := a + b
    fmt.Println("The sum is:", sum)

    // 练习3：计算两个浮点数的乘积
    x := 3.14
    y := 2.5
    product := x * y
    fmt.Println("The product is:", product)

    // 练习4：使用循环打印出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：判断一个数是否为奇数
    num := 9
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习6：使用数组存储一组整数并打印出来
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 练习7：使用切片存储一组整数并打印出来
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    // 练习8：使用map存储一组键值对并打印出来
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    fmt.Println(m)
}