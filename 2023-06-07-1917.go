package main

import "fmt"

func main() {
    // 1. 打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。
    for i := 100; i < 1000; i++ {
        a := i / 100     // 百位数
        b := (i / 10) % 10  // 十位数
        c := i % 10    // 个位数
        if i == a*a*a+b*b*b+c*c*c {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()

    // 2. 编写一个程序，找出 10000 以内的所有完数（即因子之和等于它自身的那些数）。
    for i := 1; i <= 10000; i++ {
        sum := 0   // 因子之和
        for j := 1; j < i; j++ {
            if i%j == 0 {
                sum += j
            }
        }
        if sum == i {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()

    // 3. 写出一个程序，接受一个参数，可以输出 0 到 参数指定的整数内的所有奇数。
    var n int
    fmt.Print("请输入一个整数: ")
    fmt.Scan(&n)   // 输入参数
    for i := 1; i <= n; i += 2 {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 4. 编写一个程序，找出两个正整数 a 和 b 最大的公约数。
    var a, b int
    fmt.Print("请输入两个正整数，以空格分隔: ")
    fmt.Scanf("%d %d", &a, &b)
    for b != 0 {
        a, b = b, a%b
    }
    fmt.Printf("它们的最大公约数是 %d\n", a)

    // 5. 编写一个程序，找出两个正整数 a 和 b 最小的公倍数。
    fmt.Print("请输入两个正整数，以空格分隔: ")
    fmt.Scanf("%d %d", &a, &b)
    var mul int = a
    for mul%b != 0 {
        mul += a
    }
    fmt.Printf("它们的最小公倍数是 %d\n", mul)
}