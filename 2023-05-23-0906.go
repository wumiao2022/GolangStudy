package main

import "fmt"

func main() {
	// 练习1：打印1到100的整数，但是3的倍数打印"Fizz"，5的倍数打印"Buzz"，既是3又是5的倍数打印"FizzBuzz"
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

	// 练习2：把一个字符串反转过来
	str := "hello world"
	reversedStr := ""
	for _, c := range str {
		reversedStr = string(c) + reversedStr
	}
	fmt.Println(reversedStr)

	// 练习3：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	
	// 练习4：求斐波那契数列的第n项
	n := 10
	first, second := 0, 1
	for i := 0; i < n; i++ {
		first, second = second, first+second
	}
	fmt.Println(first)
}