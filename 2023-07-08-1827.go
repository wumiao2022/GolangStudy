package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10的数值
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：判断一个数是奇数还是偶数
    num := 7
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习4：定义一个切片并打印其中的元素
    slice := []int{1, 2, 3, 4, 5}
    for _, value := range slice {
        fmt.Println(value)
    }

    // 练习5：定义一个Map并打印其中的键值对
    m := map[string]int{
        "apple":  1,
        "orange": 2,
        "banana": 3,
    }
    for key, value := range m {
        fmt.Println(key, value)
    }
}
```