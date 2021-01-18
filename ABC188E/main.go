package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

type Item struct {
	price    int
	minprice int
	value    int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	//fmt.Println(a)
	r := make(map[int][]int)
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		r[y] = append(r[y], x)
	}
	//fmt.Println(r)
	dp := make([]Item, n+1)
	dp[1] = Item{a[1], a[1], 0}
	for i := 2; i <= n; i++ {
		minprice := 1000000001
		for _, v := range r[i] {
			minprice = Min(minprice, Min(dp[v].price, dp[v].minprice))
		}
		dp[i] = Item{a[i], minprice, a[i] - minprice}
	}
	/*
		for _, v := range dp {
			fmt.Println(v)
		}
	*/
	ans := -1000000001

	// 別の町からたどり着ける道がある町だけ評価する
	for k, _ := range r {
		ans = Max(ans, dp[k].value)
	}
	fmt.Println(ans)
}

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
