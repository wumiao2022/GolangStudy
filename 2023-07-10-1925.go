package main

import "fmt"

func main() {
	// 练习1: 交换两个变量的值
	a, b := 10, 20
	a, b = b, a
	fmt.Println(a, b)

	// 练习2: 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 练习3: 实现斐波那契数列
	n := 10
	var fibonacci []int
	fibonacci = append(fibonacci, 0, 1)
	for i := 2; i <= n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// 练习4: 在切片中移除指定位置的元素
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}