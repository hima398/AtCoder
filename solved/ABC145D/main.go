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

	const p = 1000000007
	x, y := nextInt(), nextInt()
	if (x+y)%3 != 0 {
		fmt.Println(0)
		return
	}
	m := (2*x - y) / 3
	n := (2*y - x) / 3

	if m < 0 || n < 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(ModCombination(m+n, n, p))
	//fmt.Println(Combination(m+n, n) % p)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func ModCombination(N, K, p int) int {
	fac := make([]int, 1000001)
	finv := make([]int, 1000001)
	inv := make([]int, 1000001)
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < 1000001; i++ {
		fac[i] = fac[i-1] * i % p
		inv[i] = p - inv[p%i]*(p/i)%p
		finv[i] = finv[i-1] * inv[i] % p
	}
	if N < K {
		return 0
	}
	if N < 0 || K < 0 {
		return 0
	}
	return fac[N] * (finv[K] * finv[N-K] % p) % p
}
