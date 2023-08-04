package main

import "fmt"

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算1+2的结果并输出
    sum := 1 + 2
    fmt.Println("1 + 2 =", sum)

    // 练习3：定义一个整型变量num并赋值为10，输出num的值
    num := 10
    fmt.Println("num =", num)

    // 练习4：定义一个切片slices，包含元素1、2、3，并输出切片的长度和容量
    slices := []int{1, 2, 3}
    fmt.Println("slices length =", len(slices))
    fmt.Println("slices capacity =", cap(slices))

    // 练习5：定义一个map，包含键值对age: 20、name: "John"，并输出map的内容
    person := map[string]interface{}{"age": 20, "name": "John"}
    fmt.Println("person =", person)
}