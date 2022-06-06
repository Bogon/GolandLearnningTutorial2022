package main

import "fmt"

// 运算符
func main() {
	var (
		a = 3
		b = 4
	)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	/// 自增、自减运算符只能在变量的右边 等价于：a = a + 1, b = b -1
	a++
	b--

	// 关系运算符
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)

	// 逻辑运算符
	if a == 2 || b == 3 {
		fmt.Println("a == 2 || b == 3")
	}

	if a == 2 && b == 3 {
		fmt.Println("a == 2 && b == 3")
	}

	if !(a == 2 && b == 3) {
		fmt.Println("!(a == 2 && b == 3)")
	}

	// 位运算 针对二进制数
	// 5 的二进制位：101
	// 2 的二进制位：10

	// & 运算符：按位与 都为1 才为1
	fmt.Println(5 & 2)

	// | 按位或 有一个为1 就是1
	fmt.Println(5 | 2)

	// ^ 异或 参与运算的数字两位不一样为1
	fmt.Println(5 ^ 2)

	// << 左移运算符
	fmt.Println(5 << 1)

	// << 右移运算符
	fmt.Println(5 >> 1)

	// 赋值运算符
	var x int
	x = 10
	x += 1
	x -= 1
	x *= 2
	x /= 2
	x %= 2
	x <<= 2
	x >>= 2
	x &= 2
	x |= 2
	x ^= 2
}
