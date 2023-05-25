package main

import (
	"fmt"
)

func main() {
	// 练习1：打印数字1到100，并在能被3整除时输出"Fizz"，在能被5整除时输出"Buzz"，在能同时被3和5整除时输出"FizzBuzz"
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

	// 练习2：将一个字符串进行反转输出
	s := "hello, world"
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	fmt.Println(string(rs))
	
	// 练习3：求一个整数数组中的最大值和最小值
	numbers := []int{10, 23, 45, 67, 89, 1, 4, -2, -100}
	max, min := numbers[0], numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Printf("Max: %d, Min: %d\n", max, min)
}