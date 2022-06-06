package main

import "fmt"

var (
	coins        = 5000
	users        = []string{"Mattew", "Sarah", "Augustus", "Heidi", "Emillie", "Patter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func dispatchCoins() (left int) {
	// 1、一次拿到每个人的名字
	for _, name := range users {
		// 2、拿到一个人的名字，按照规则分金币
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--

			case 'i', 'I':
				distribution[name] += 2
				coins -= 2

			case 'o', 'O':
				distribution[name] += 3
				coins -= 3

			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	// 2.1 每个人分得的金币应该放到 distribution中
	// 2.2 计算剩余多少金币
	// 3、最后返回剩余金币数
	left = coins
	return
}

func main() {
	ret := dispatchCoins()
	fmt.Println("剩下：", ret)
}
