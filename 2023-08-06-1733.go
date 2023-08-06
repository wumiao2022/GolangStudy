package main

import "fmt"

// 练习1: 打印Hello, World!
func main() {
	fmt.Println("Hello, World!")
}

// 练习2: 定义一个函数，接收两个参数并返回它们的和
func add(a, b int) int {
	return a + b
}

// 练习3: 定义一个结构体Person，包含姓名和年龄字段，并实例化一个Person对象
type Person struct {
	name string
	age  int
}

// 练习4: 实现一个接口，并调用该接口的方法
type Greeting interface {
	greet()
}

func (p Person) greet() {
	fmt.Println("Hello, my name is", p.name)
}

func main() {
	p := Person{"Alice", 25}
	p.greet()
}

// 练习5: 使用并发技术实现并行计算素数的数量
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- countPrimes(1, 100000)
	}()
	go func() {
		ch <- countPrimes(100001, 200000)
	}()
	total := <-ch + <-ch
	fmt.Println("Total primes found:", total)
}