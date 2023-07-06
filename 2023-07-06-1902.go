package main

import "fmt"

func main() {
    // 练习1：输出Hello World
    fmt.Println("Hello World")

    // 练习2：定义并输出整型变量i的值
    var i int
    i = 10
    fmt.Println(i)

    // 练习3：定义并输出浮点型变量f的值
    var f float64
    f = 3.14
    fmt.Println(f)

    // 练习4：定义并输出字符串变量s的值
    var s string
    s = "Hello Golang"
    fmt.Println(s)

    // 练习5：使用带有指定格式的Printf函数输出名字和年龄
    name := "Alice"
    age := 25
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    // 练习6：定义数组arr，并输出其长度和第一个元素的值
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Length of arr:", len(arr))
    fmt.Println("First element of arr:", arr[0])

    // 练习7：定义切片slice，并输出其长度和第一个元素的值
    slice := arr[1:4]
    fmt.Println("Length of slice:", len(slice))
    fmt.Println("First element of slice:", slice[0])

    // 练习8：定义并输出map类型变量m的值
    m := make(map[string]int)
    m["apple"] = 5
    m["banana"] = 3
    fmt.Println("map m:", m)

    // 练习9：使用for循环输出1到10之间的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习10：使用if条件判断语句判断某个数是奇数还是偶数，并输出结果
    num := 7
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```