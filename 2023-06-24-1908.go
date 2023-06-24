package main

import "fmt"

func main() {
	// 1. 声明并初始化变量
	var a int = 10
	var b float64 = 3.14

	// 2. 算术运算
	fmt.Println(a + 5)      // 15
	fmt.Println(b - 1.0)    // 2.14
	fmt.Println(a * 2)      // 20
	fmt.Println(b / 2.0)    // 1.57
	fmt.Println(a % 3)      // 1
	fmt.Println(a &^ 3)     // 8
	fmt.Println(a << 1)     // 20
	fmt.Println(a >> 1)     // 5
	fmt.Println(a & 1)      // 0
	fmt.Println(a | 1)      // 11
	fmt.Println(a ^ 1)      // 11
	fmt.Println(^a)         // -11
	fmt.Println(^a + 1)     // -10
	fmt.Println(1 << 10)    // 1024
	fmt.Println(2 ^ 3 ^ 2)  // 3
	fmt.Println(2 ^ 2 ^ 3)  // 3

	// 3. 赋值运算
	a = 20
	b += 1.0
	fmt.Println(a, b)

	// 4. 比较运算
	fmt.Println(a == 20)    // true
	fmt.Println(b != 3.14)  // true
	fmt.Println(a > 10)     // true
	fmt.Println(b < 4.0)    // true
	fmt.Println(a <= 20)    // true
	fmt.Println(b >= 3.14)  // true

	// 5. 逻辑运算
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(!true)          // false
	fmt.Println(!(a > 10))      // false
	fmt.Println(!(b < 4.0))     // false

	// 6. 位运算
	fmt.Println(1 << 1.5)  // 2
	fmt.Println(1.5 << 1)  // invalid operation: 1.5 << 1 (shift count type float64, must be integer)

	// 7. 数组
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)            // [1 2 3]
	fmt.Println(arr[0])         // 1
	fmt.Println(len(arr))       // 3
	fmt.Println(arr[:2])        // [1 2]
	fmt.Println(cap(arr))       // 3
	fmt.Printf("%T\n", arr)     // [3]int
	fmt.Printf("%#v\n", arr)    // [3]int{1, 2, 3}
	fmt.Println(&arr[0] == &arr[2])  // false

	// 8. 切片
	slice1 := []int{1,2,3,4,5}
	fmt.Println(slice1)                // [1 2 3 4 5]
	fmt.Println(len(slice1))           // 5
	fmt.Println(cap(slice1))           // 5
	fmt.Printf("%T\n", slice1)         // []int
	fmt.Printf("%#v\n", slice1)        // []int{1, 2, 3, 4, 5}

	slice2 := slice1[1:3]
	fmt.Println(slice2)                // [2 3]
	fmt.Println(len(slice2))           // 2
	fmt.Println(cap(slice2))           // 4
	fmt.Printf("%T\n", slice2)         // []int
	fmt.Printf("%#v\n", slice2)        // []int{2, 3}

	slice3 := make([]int, 3, 5)
	fmt.Println(slice3)                // [0 0 0]
	fmt.Println(len(slice3))           // 3
	fmt.Println(cap(slice3))           // 5
	fmt.Printf("%T\n", slice3)         // []int
	fmt.Printf("%#v\n", slice3)        // []int{0, 0, 0}

	slice4 := append(slice3, 1, 2, 3, 4)
	fmt.Println(slice4)                // [0 0 0 1 2 3 4]
	fmt.Println(len(slice4))           // 7
	fmt.Println(cap(slice4))           // 10
	fmt.Printf("%T\n", slice4)         // []int
	fmt.Printf("%#v\n", slice4)        // []int{0, 0, 0, 1, 2, 3, 4}

	// 9. map
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)              // map[one:1 two:2]
	fmt.Println(m["one"])       // 1
	fmt.Println(m["three"])     // 0

	delete(m, "two")
	fmt.Println(m)              // map[one:1]

	// 10. if语句
	if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a <= 10")
	}

	// 11. for循环
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}

	// 12. select语句
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1:", num1)
	case num2 := <-ch2:
		fmt.Println("ch2:", num2)
	}
}