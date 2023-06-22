package main

import (
    "fmt"
)

func main() {
    // 1. 交换两个变量的值
    a, b := 10, 20
    a, b = b, a
    fmt.Println(a, b)

    // 2. 循环打印出1到100中能被3整除的数字
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }

    // 3. 定义一个匿名结构体，并输出其中的字段值
    person := struct {
        name string
        age int
    }{"Tom", 20}
    fmt.Println(person.name, person.age)

    // 4. 实现一个函数，可以接收任意数量的int类型参数，并返回它们的和
    fmt.Println(sum(1, 2, 3, 4))

    // 5. 实现一个函数，可以接收一个字符串作为参数，并返回这个字符串中每个字符出现的次数
    fmt.Println(countChar("hello"))
}

func sum(nums ...int) int {
    result := 0
    for _, num := range nums {
        result += num
    }
    return result
}

func countChar(str string) map[rune]int {
    result := make(map[rune]int)
    for _, char := range str {
        result[char]++
    }
    return result
}