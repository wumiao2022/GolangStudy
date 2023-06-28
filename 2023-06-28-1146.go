package main

import "fmt"

func main() {
    // 1. 打印Hello, World!
    fmt.Println("Hello, World!")

    // 2. 定义一个字符串变量，赋值为"GoLang"
    var str string = "GoLang"
	
    // 3. 判断字符串变量是否为空
    if str == "" {
        fmt.Println("字符串为空")
    } else {
        fmt.Println("字符串不为空")
    }
    
    // 4. 判断一个数是奇数还是偶数
    num := 10
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }
    
    // 5. 使用for循环打印1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 6. 定义一个切片，包含1到5的整数
    slice := []int{1, 2, 3, 4, 5}
    
    // 7. 使用for循环遍历切片并打印每个元素
    for _, value := range slice {
        fmt.Println(value)
    }
    
    // 8. 定义一个函数，接受两个参数并返回它们的和
    func add(a, b int) int {
        return a + b
    }
    
    // 9. 调用上面定义的函数，传递参数为2和3，打印返回值
    sum := add(2, 3)
    fmt.Println(sum)
    
    // 10. 定义一个结构体类型Person，包含name和age字段
    type Person struct {
        name string
        age  int
    }
    
    // 11. 创建一个Person类型的变量并初始化
    p := Person{
        name: "Tom",
        age:  25,
    }
    
    // 12. 打印Person变量的name和age字段
    fmt.Println(p.name, p.age)
}