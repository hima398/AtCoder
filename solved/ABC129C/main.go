package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	ma := make(map[int]bool)
	for i := 0; i < m; i++ {
		a := nextInt()
		ma[a] = true
	}
	const p = 1e9 + 7
	//fmt.Println(p)

	dp := make([]int, n+1) // 0 ... n
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if ma[i] {
			continue // dp[i] = 0
		}
		for j := i - 1; j >= i-2 && j >= 0; j-- {
			if ma[j] {
				continue
			}
			dp[i] += dp[j]
			dp[i] %= p
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[n])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
