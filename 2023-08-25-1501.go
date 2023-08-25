package main

import "fmt"

func main() {
	// 练习1：输出"Hello, World!"到控制台
	fmt.Println("Hello, World!")

	// 练习2：打印1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习3：计算1到100之间所有数字的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：反转字符串
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)
}

注意：以上代码是练习的示例代码，不包含任何解释。对于理解代码的具体功能和细节，你可以查阅相关文档或是其他教程。