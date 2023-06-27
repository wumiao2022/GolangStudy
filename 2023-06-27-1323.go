package main

import "fmt"

func main() {
	// 练习1：输出Hello world！
	fmt.Println("Hello world!")

	// 练习2：输出1~10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：输出1~100中的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是素数\n", num)
	} else {
		fmt.Printf("%d不是素数\n", num)
	}

	// 练习5：将一个字符串反转输出
	str := "hello"
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	fmt.Println(res)
}