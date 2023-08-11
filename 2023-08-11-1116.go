package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：声明一个变量并赋值，然后打印该变量的值
    var num int = 10
    fmt.Println(num)

    // 练习3：声明一个数组并初始化，然后打印数组的元素
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 练习4：使用for循环打印1到10的所有奇数
    for i := 1; i <= 10; i += 2 {
        fmt.Println(i)
    }

    // 练习5：使用if-else语句判断一个数是奇数还是偶数，并打印结果
    num2 := 15
    if num2%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }
}