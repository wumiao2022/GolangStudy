package main

import "fmt"

// 练习1：使用for循环打印出1到10的数字
func forLoop() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}

// 练习2：定义一个函数，将两个整数相加并返回结果
func add(a, b int) int {
    return a + b
}

// 练习3：编写一个函数，判断一个字符串是否是回文字符串
func isPalindrome(str string) bool {
    for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-1-i] {
            return false
        }
    }
    return true
}

// 练习4：定义结构体"Person"，包含姓名和年龄两个字段
type Person struct {
    name string
    age  int
}

// 练习5：实现结构体方法，用于打印Person信息
func (p Person) PrintInfo() {
    fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func main() {
    // 练习1
    forLoop()

    // 练习2
    sum := add(3, 4)
    fmt.Println("Sum:", sum)

    // 练习3
    isPalin := isPalindrome("level")
    fmt.Println("Is Palindrome:", isPalin)

    // 练习4
    person := Person{name: "Alice", age: 25}
    person.PrintInfo()
}