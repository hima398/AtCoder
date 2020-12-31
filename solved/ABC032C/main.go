package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}

type Pos struct {
	X int
	Y int
}

func SolveStupidly(N, K int) int {
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = nextInt()
	}
	ans := 0
	for i := 0; i < N; i++ {
		total := 1
		for j := i; j < N; j++ {
			total *= s[j]
			if total > K {
				continue
			}
			ans = Max(ans, j-i+1)
		}
	}
	return ans
}

func Solve(n, k int) int {
	s := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		s[i] = nextInt()
		if s[i] == 0 {
			return n
		}
	}
	r := 0
	total := 1
	for l := 0; l < n; l++ {
		for r < n && total*s[r] <= k {
			total *= s[r]
			r++
		}
		ans = Max(ans, r-l)
		if l == r {
			r++
		} else {
			// 要素に0が無いのは入力時点で保証済み
			total /= s[l]
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	//fmt.Println(SolveStupidly(n, k))
	fmt.Println(Solve(n, k))
}
