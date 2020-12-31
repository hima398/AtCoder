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

func SolveStupidly(n int) int {
	if n == 1 {
		return 1
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	ans := 0
	for i := 0; i < n; i++ {
		m := make(map[int]bool)
		//m[a[i]] = true
		for j := i; j < n; j++ {
			if m[a[j]] {
				//fmt.Println(m)
				ans = Max(ans, j-i)
				break
			} else if j == n-1 {
				ans = Max(ans, j-i+1)
				break
			}
			m[a[j]] = true
		}
	}

	return ans
}

func SolveTwoPoint(n int) int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	r := 0
	ans := 0
	m := make(map[int]bool)
	for l := 0; l < n; l++ {
		for r < n {
			if m[a[r]] {
				break
			}
			m[a[r]] = true
			r++
		}
		ans = Max(ans, r-l)
		delete(m, a[l])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	//fmt.Println(SolveStupidly(n))
	fmt.Println(SolveTwoPoint(n))
}
