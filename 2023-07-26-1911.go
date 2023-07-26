package main

import "fmt"

func main() {
    // 实现一个函数，判断一个数字是否为偶数并输出结果
    num := 10
    if num%2 == 0 {
        fmt.Println("是偶数")
    } else {
        fmt.Println("不是偶数")
    }
}