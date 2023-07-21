package main

import "fmt"

func main() {
	// 示例1: 变量声明与初始化
	var num1 int    // 声明一个int类型的变量num1
	num1 = 10       // 给变量num1赋值为10

	var str string = "Hello World"   // 声明一个string类型的变量str并初始化为"Hello World"
	num2 := 20     // 声明一个int类型的变量num2，并使用简短声明初始化为20

	fmt.Println("num1:", num1)
	fmt.Println("str:", str)
	fmt.Println("num2:", num2)

	// 示例2: 数组遍历
	arr := [5]int{1, 2, 3, 4, 5}    // 声明并初始化一个长度为5的int数组
	for i := 0; i < len(arr); i++ { // 循环遍历数组元素并打印
		fmt.Println("arr["+fmt.Sprintf("%d", i)+"]:", arr[i])
	}

	// 示例3: 切片操作
	slice := arr[1:4]     // 创建一个从数组arr索引1到3的切片
	fmt.Println("slice:", slice)

	// 示例4: 切片追加元素
	slice = append(slice, 6)     // 在切片末尾追加元素6
	fmt.Println("slice after appending 6:", slice)
}