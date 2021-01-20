package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(s, l, r string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i <= n; i++ {
		dp[i] = Min(dp[i], dp[i-1]+1)
		if s[i-2:i] == l || s[i-2:i] == r {
			dp[i] = Min(dp[i], dp[i-2]+1)
		}
	}
	return dp[n]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const button = "ABXY"
	var ls []string
	var rs []string
	for _, l1 := range button {
		for _, l2 := range button {
			ls = append(ls, string(l1)+string(l2))
		}
	}
	for _, r1 := range button {
		for _, r2 := range button {
			rs = append(rs, string(r1)+string(r2))
		}
	}

	n := nextInt()
	s := nextString()

	ans := n
	for _, l := range ls {
		for _, r := range rs {
			ans = Min(ans, Solve(s, l, r))
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
