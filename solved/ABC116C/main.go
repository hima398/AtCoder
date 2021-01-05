package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func SolveLoop(h []int) int {
	prev := 0
	ans := 0
	for _, v := range h {
		ans += Max(v-prev, 0)
		prev = v
	}
	return ans
}

func SolveRecursive(h []int) int {
	//fmt.Println(h)
	if len(h) == 0 {
		return 0
	}
	if len(h) == 1 {
		return h[0]
	}
	min := h[0]
	// 最小を求めるついでに全て同じかを調べておく
	IsSame := true
	for i := 1; i < len(h); i++ {
		min = Min(min, h[i])
		IsSame = IsSame && (h[i] == h[i-1])
	}
	if IsSame {
		return h[0]
	}
	ans := min
	m := -1
	for i := 0; i < len(h); i++ {
		h[i] -= min
		if m == -1 && h[i] == 0 {
			m = i
		}
	}

	if m >= 0 {
		// 途中に0がない
		ans += SolveRecursive(h[:m]) + SolveRecursive(h[m+1:])
		return ans
	} else {
		return ans
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	h := make([]int, n)

	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}
	//fmt.Println(SolveRecursive(h))
	fmt.Println(SolveLoop(h))

}
