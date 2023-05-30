package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：定义一个整型变量，赋值为10，打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量，赋值为"Hello Golang"，打印该变量的值
	var str string = "Hello Golang"
	fmt.Println(str)

	// 练习4：声明一个长度为3的字符串数组，分别赋值为"apple"、"banana"、"orange"，并遍历打印出来
	var fruits [3]string = [3]string{"apple", "banana", "orange"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// 练习5：定义一个切片slice，赋值为1、2、3、4、5，添加一个元素6，再删除一个元素4，最后打印该slice
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	slice = append(slice[:3], slice[4:]...)
	fmt.Println(slice)

	// 练习6：定义一个map，用于存储学生的姓名和年龄，分别存储"Tom":18、"Jerry":17、"Kate":19，最后遍历打印出来
	studentMap := map[string]int{"Tom": 18, "Jerry": 17, "Kate": 19}
	for name, age := range studentMap {
		fmt.Println(name, age)
	}
}