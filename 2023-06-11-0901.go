package main

import (
    "fmt"
)

func main() {
    // 练习1：计算1到10的和
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("1到10的和是：", sum)
    
    // 练习2：使用for循环打印1到100中所有能被3整除的数字
    for i := 1; i <= 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
    
    // 练习3：使用for循环和if语句判断一个整数是否为素数
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
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
    
    // 练习4：使用switch语句判断用户输入的数字代表的星期几，并输出相应的信息
    var day int
    fmt.Print("请输入数字代表的星期几：")
    fmt.Scan(&day)
    switch day {
    case 1:
        fmt.Println("星期一")
    case 2:
        fmt.Println("星期二")
    case 3:
        fmt.Println("星期三")
    case 4:
        fmt.Println("星期四")
    case 5:
        fmt.Println("星期五")
    case 6:
        fmt.Println("星期六")
    case 7:
        fmt.Println("星期日")
    default:
        fmt.Println("输入有误")
    }
    
}