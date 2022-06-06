package main

import "fmt"

// 递归：自己调用自己
// 递归一定要有一个推出的条件，
// 递归适合解决相同问题，问题的规模越来越小的场景
// 永远不要高估自己！

// 计算 n 的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶面试题
// n个台阶，一次可以走一步，也可走两步，有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return taijie(n-1) + taijie(n-2)
}

func main() {
	// ret := f(7)
	ret := taijie(9)
	fmt.Println(ret)
}
