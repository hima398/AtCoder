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

func MinFloat(x, y float64) float64 {
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

func MaxFloat(x, y float64) float64 {
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

func ExpectedValue(n int) float64 {
	if n == 1 {
		return 1.0
	}
	fn := float64(n)
	return (fn + 1) * fn / (2 * fn)
}

func SolveSlidingWindow(n, k int) float64 {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}
	sum := 0.0
	ans := 0.0
	for i := 0; i < k-1; i++ {
		sum += ExpectedValue(p[i])
	}
	for l := 0; l <= n-k; l++ {
		r := l + k
		sum += ExpectedValue(p[r-1])
		//fmt.Println(l, r)
		ans = MaxFloat(ans, sum)
		sum -= ExpectedValue(p[l])
	}
	return ans
}

func SolvePrefixSum(n, k int) float64 {
	s := make([]float64, n+1)
	for i := 0; i < n; i++ {
		p := nextInt()
		// piのサイコロの期待値
		ex := ExpectedValue(p)
		// piのサイコロの期待値の累積和
		s[i+1] = s[i] + ex
	}
	ans := 0.0
	for l := 0; l <= n-k; l++ {
		r := l + k
		//fmt.Println(l, r)
		ans = MaxFloat(ans, s[r]-s[l])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	fmt.Println(SolvePrefixSum(n, k))
	//fmt.Println(SolveSlidingWindow(n, k))
}
