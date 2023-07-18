package main

import (
	"fmt"
)

func main() {
	Exercise1()
}

// Exercise1 请编写一个程序，打印出从1到100的所有奇数。
func Exercise1() {
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}
}