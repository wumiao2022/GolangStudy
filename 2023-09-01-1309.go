package main

import "fmt"

func main() {
    // 练习1: 输出 "Hello, Golang!"
    fmt.Println("Hello, Golang!")

    // 练习2: 声明一个整数变量并赋值为10，然后输出该变量的值
    num := 10
    fmt.Println(num)

    // 练习3: 声明一个浮点数变量并赋值为3.14，然后输出该变量的值
    floatNum := 3.14
    fmt.Println(floatNum)

    // 练习4: 声明一个字符串变量并赋值为"Hello, World!"，然后输出该变量的值
    str := "Hello, World!"
    fmt.Println(str)

    // 练习5: 声明一个布尔型变量并赋值为true，然后输出该变量的值
    boolVar := true
    fmt.Println(boolVar)

    // 练习6: 声明一个整数数组并初始化为1,2,3,4,5，然后输出该数组的值
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 练习7: 声明一个切片变量并初始化为1,2,3,4,5，然后输出该切片的值
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    // 练习8: 使用for循环输出1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习9: 使用if语句判断一个数是否为偶数，如果是则输出 "偶数"，否则输出 "奇数"
    num = 7
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习10: 声明一个结构体类型Person，包含name和age两个字段，然后创建一个Person变量并赋值，
    // 然后输出该变量的字段值
    type Person struct {
        name string
        age  int
    }
    person := Person{name: "Alice", age: 20}
    fmt.Println(person.name, person.age)
}
