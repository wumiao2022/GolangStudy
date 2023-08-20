package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个变量并赋值，然后打印该变量的值
	num := 10
	fmt.Println(num)

	// 3. 使用if语句判断一个数是否为偶数，如果是偶数则打印"偶数"，否则打印"奇数"
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 4. 使用for循环打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 定义一个slice，并使用append函数向其中追加元素，然后打印slice的长度和所有元素
	slice := []int{}
	slice = append(slice, 1, 2, 3)
	fmt.Println(len(slice))
	fmt.Println(slice)

	// 6. 定义一个map，并向其中插入键值对，然后根据键查找值并打印
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	fmt.Println(m["A"])

	// 7. 使用defer关键字延迟执行一个函数，在函数返回前打印一条消息
	defer fmt.Println("deferred function")

	// 8. 使用panic关键字触发一个panic，并在defer函数中使用recover函数捕获并打印panic的值
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()
	panic("panic occurred")
}