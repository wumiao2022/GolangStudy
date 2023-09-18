package main

import "fmt"

// 练习 #1：打印 Hello, World!
func exercise1() {
  fmt.Println("Hello, World!")
}

// 练习 #2：计算两个整数的和并打印结果
func exercise2() {
  a := 10
  b := 20
  sum := a + b
  fmt.Println(sum)
}

// 练习 #3：判断一个整数是否为偶数并打印结果
func exercise3() {
  num := 15
  if num % 2 == 0 {
    fmt.Println("该整数是偶数")
  } else {
    fmt.Println("该整数是奇数")
  }
}

// 练习 #4：循环打印数字 1 到 5
func exercise4() {
  for i := 1; i <= 5; i++ {
    fmt.Println(i)
  }
}

// 练习 #5：定义一个结构体类型并创建一个结构体实例
type Person struct {
  Name string
  Age  int
}

func exercise5() {
  p := Person{Name: "Alice", Age: 25}
  fmt.Println(p)
}

func main() {
  exercise1()
  exercise2()
  exercise3()
  exercise4()
  exercise5()
}
