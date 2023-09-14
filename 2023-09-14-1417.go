package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Println("Current time:", time.Now())

	// 定义一个整型数组
	numbers := []int{1, 2, 3, 4, 5}

	// 计算数组元素之和
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of numbers:", sum)

	// 打印1到100的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 判断一个整数是否为素数
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	fmt.Println("Is 7 prime?", isPrime(7))
	fmt.Println("Is 10 prime?", isPrime(10))

	// 定义一个结构体
	type Person struct {
		Name string
		Age  int
	}

	// 创建一个Person实例
	person := Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(person)

	// 定义一个接口
	type Animal interface {
		Sound() string
	}

	// 定义一个实现Animal接口的结构体
	type Cat struct{}

	// 实现接口方法
	func (c Cat) Sound() string {
		return "Meow"
	}

	// 创建一个Cat实例并调用接口方法
	cat := Cat{}
	fmt.Println(cat.Sound())
}