package main

import (
	"fmt"
)

func main() {
	// 练习题 1：输出1-100的所有偶数
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习题 2：输出一个切片的所有元素
	slice := []int{1, 2, 3, 4, 5}
	for _, v := range slice {
		fmt.Println(v)
	}

	// 练习题 3：计算斐波那契数列前20个数字
	fib1, fib2 := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println(fib2)
		temp := fib2
		fib2 += fib1
		fib1 = temp
	}

	// 练习题 4：输出一个map的所有键值对
	m := map[string]string{"name": "张三", "age": "18", "gender": "男"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}