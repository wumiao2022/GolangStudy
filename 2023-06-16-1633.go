package main

import (
	"fmt"
)

func main() {
	// 练习：打印1-100的数字，如果是3的倍数打印“Fizz”，如果是5的倍数打印“Buzz”，如果既是3的倍数又是5的倍数打印“FizzBuzz”
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}