package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印出1到100的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2: 求1到100的所有数字的和，输出总和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("总和:", sum)

	// 练习3: 判断一个数是否为素数（只能被1和自身整除的数）
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
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

	// 练习4: 交换两个变量的值
	a := 5
	b := 10
	fmt.Println("交换前：a =", a, ", b =", b)
	temp := a
	a = b
	b = temp
	fmt.Println("交换后：a =", a, ", b =", b)
}
```