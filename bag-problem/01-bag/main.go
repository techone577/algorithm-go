package main

import "fmt"

// 递推公式:
// B(k, w) = B(k-1, w) (不选取第 k 件物品) or max{ B(k - 1, w - wk) + Vk (背包剩余容量够容纳第 k 件), B(k-1, w)(不够容纳第 k 件) } (选取第 k 件物品)
// B(k, w): 选取前 k 件物品， 背包剩余容量为 w 情况下的最大价值
//

type thing struct {
	W int
	V int
}

var things = []thing{
	{2, 3},
	{3, 4},
	{4, 5},
	{5, 8},
	{9, 10},
}

func B(k, w int) int {
	if w <= 0 {
		return 0
	}
	if k == 0 {
		if things[k].W <= w {
			return things[k].V
		} else {
			return 0
		}
	}
	// not pick k
	notPick := func() int { return B(k-1, w) }
	// pick k
	// can pick / can not pick / not enough w
	pick := func() int { return max(B(k-1, w-things[k].W)+things[k].V, B(k-1, w)) }

	if things[k].W <= w {
		return max(pick(), notPick())
	} else {
		return notPick()
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(B(4, 20))
}
