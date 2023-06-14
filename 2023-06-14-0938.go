package main

import "fmt"

func main() {
	// 练习1：输出“Hello, World!”
	fmt.Println("Hello, World!")

	// 练习2：输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：输出1到100中，能被3整除，但不能被5整除的数字
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}

	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习5：将一个字符串反转输出
	str := "hello, world"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)
}