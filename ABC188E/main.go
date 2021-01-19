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

func Solve(n, m int) int {
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
	return ans
}

func SolveCommentary1(n, m int) int {
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	//fmt.Println(a)
	r := make(map[int][]int)
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		r[x] = append(r[x], y)
	}
	//fmt.Println(r)
	// dp[i] 街iにたどり着くまでに購入できる最小の金額
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1000000001
	}
	for i := 1; i <= n; i++ {
		for _, v := range r[i] {
			dp[v] = Min(dp[v], dp[i])
			dp[v] = Min(dp[v], a[i])
		}
	}
	//fmt.Println(dp)
	ans := -1000000001
	for i := 1; i <= n; i++ {
		ans = Max(ans, a[i]-dp[i])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	//fmt.Println(Solve(n, m))
	fmt.Println(SolveCommentary1(n, m))
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
