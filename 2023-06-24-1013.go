package main

import (
	"fmt"
)

func main() {
	// 1. 将字符串转换为整型
	str := "123"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)

	// 2. 将整型转换为字符串
	num = 456
	str = strconv.Itoa(num)
	fmt.Println(str)

	// 3. 将字符串转换为浮点型
	str = "1.23"
	f, _ := strconv.ParseFloat(str, 64)
	fmt.Println(f)

	// 4. 将浮点型转换为字符串
	f = 4.56
	str = strconv.FormatFloat(f, 'g', -1, 64)
	fmt.Println(str)

	// 5. 取正负号
	num = -7
	fmt.Println(math.Signbit(float64(num)))

	// 6. 求绝对值
	num = -10
	fmt.Println(math.Abs(float64(num)))

	// 7. 向上取整
	f = 1.4
	fmt.Println(math.Ceil(f))

	// 8. 向下取整
	f = 1.6
	fmt.Println(math.Floor(f))

	// 9. 求最大值
	num1 := 5
	num2 := 3
	fmt.Println(math.Max(float64(num1), float64(num2)))

	// 10. 求最小值
	fmt.Println(math.Min(float64(num1), float64(num2)))
}