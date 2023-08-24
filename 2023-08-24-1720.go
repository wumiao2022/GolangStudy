package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：交换两个变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 练习3：计算给定数组的平均值
	numbers := []float64{1, 2, 3, 4, 5}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))
	fmt.Println("Average =", average)

	// 练习4：判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(num, "is prime:", isPrime)

	// 练习5：将英文名字全大写，并在前面加上"Mr."
	name := "john doe"
	name = "Mr. " + name
	name = strings.ToUpper(name)
	fmt.Println("Formatted name:", name)
}