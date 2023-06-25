package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 4. 输出10的阶乘
	factorial := 1
	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

	// 5. 打印1到100的数字，如果是3的倍数则输出Fizz，如果是5的倍数则输出Buzz，如果同时是3和5的倍数则输出FizzBuzz
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