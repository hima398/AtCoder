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

func Eval(l, r int, s, x []int) bool {

	sum := s[r] - s[l-1]
	xor := x[r] ^ x[l-1]
	//fmt.Printf("sum = %d, xor = %d\n", sum, xor)
	return sum == xor
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n+1)
	s := make([]int, n+1)
	x := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		s[i] += s[i-1] + a[i]
		x[i] ^= x[i-1] ^ a[i]
	}
	ans := 0
	r := 1
	for l := 1; l <= n; l++ {
		for r <= n && Eval(l, r, s, x) {
			//fmt.Println(l, r)
			r++
		}
		ans += r - l
		if l == r {
			r++
		}
	}
	fmt.Println(ans)
}
