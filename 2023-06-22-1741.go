package main

import (
    "fmt"
)

func main() {
    // 练习1：定义一个常量，值为100，然后输出这个常量
    const num = 100
    fmt.Println(num)

    // 练习2：定义一个变量x，值为10，然后在变量x的基础上进行加减乘除赋值运算，并输出最终的值
    var x = 10
    x += 5 // 等价于 x = x + 5
    fmt.Println(x)
    x -= 3 // 等价于 x = x - 3
    fmt.Println(x)
    x *= 2 // 等价于 x = x * 2
    fmt.Println(x)
    x /= 4 // 等价于 x = x / 4
    fmt.Println(x)

    // 练习3：定义一个长度为5的整型数组，并将数组中的元素初始化为1、2、3、4、5，然后输出这个数组
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 练习4：定义一个字符串变量，值为"Hello, World!"，然后将其中的字符'o'替换为'0'，并输出最终的字符串
    var str string = "Hello, World!"
    str = string([]rune(str)[:4]) + "0" + string([]rune(str)[5:])
    fmt.Println(str)

    // 练习5：定义一个函数，功能为计算两个整型数的和，并输出结果
    var a, b int = 10, 20
    sum := add(a, b)
    fmt.Println(sum)
}

func add(a, b int) int {
    return a + b
}