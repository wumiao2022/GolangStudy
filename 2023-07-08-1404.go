package main

import "fmt"

func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello World")

    // 练习2：定义一个变量，赋值为10，打印出该变量的值
    var num int = 10
    fmt.Println(num)

    // 练习3：定义一个切片，包含数字1到5，使用for循环打印切片的每个元素
    nums := []int{1, 2, 3, 4, 5}
    for _, num := range nums {
        fmt.Println(num)
    }

    // 练习4：定义一个函数，接收两个整数作为参数，返回它们的和
    func add(a, b int) int {
        return a + b
    }

    // 练习5：使用切片实现一个栈，包括Push和Pop方法
    type Stack struct {
        elements []int
    }

    func (s *Stack) Push(num int) {
        s.elements = append(s.elements, num)
    }

    func (s *Stack) Pop() int {
        if len(s.elements) > 0 {
            num := s.elements[len(s.elements)-1]
            s.elements = s.elements[:len(s.elements)-1]
            return num
        }
        return 0
    }

    // 创建一个栈实例
    stack := Stack{}

    // 使用Push方法向栈中添加元素
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    // 使用Pop方法从栈中取出元素并打印
    fmt.Println(stack.Pop())
    fmt.Println(stack.Pop())
    fmt.Println(stack.Pop())
}