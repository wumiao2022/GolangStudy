package main

import "fmt"

func main() {
  // 练习1：打印Hello, World!
  fmt.Println("Hello, World!")

  // 练习2：计算1+2的结果并打印
  fmt.Println(1 + 2)

  // 练习3：声明一个整数变量x并赋值为10，然后打印出来
  x := 10
  fmt.Println(x)

  // 练习4：声明一个字符串变量name并赋值为"John"，然后打印出来
  name := "John"
  fmt.Println(name)

  // 练习5：声明一个布尔型变量isTrue并赋值为true，然后打印出来
  isTrue := true
  fmt.Println(isTrue)

  // 练习6：声明一个数组nums包含5个整数，分别为1, 2, 3, 4, 5，并打印出来
  nums := [5]int{1, 2, 3, 4, 5}
  fmt.Println(nums)

  // 练习7：使用for循环打印出1到10的数字，每个数字独占一行
  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }

  // 练习8：编写一个函数add，接收两个整数参数并返回它们的和，然后调用该函数并打印结果
  fmt.Println(add(3, 4))
}

func add(a, b int) int {
  return a + b
}
