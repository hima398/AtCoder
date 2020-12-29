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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, M := nextInt(), nextInt()
	S := make([]int, M)
	T := make([]int, M)
	Table := make([]int, N+2)
	for i := 0; i < M; i++ {
		S[i] = nextInt()
		T[i] = nextInt()
		Table[S[i]]++
		Table[T[i]+1]--
	}
	fmt.Println(Table)
	for i := 0; i < N; i++ {
		Table[i+1] += Table[i]
	}
	fmt.Println(Table)
	var ans []int
	for i := 0; i < M; i++ {
		ok := true
		for j := S[i]; j <= T[i]; j++ {
			if Table[j] <= 1 {
				ok = false
				break
			}
		}
		if ok {
			ans = append(ans, i+1)
		}
	}
	wc := bufio.NewWriter(os.Stdout)
	defer wc.Flush()
	wc.WriteString(strconv.Itoa(len(ans)) + "\n")
	for _, v := range ans {
		wc.WriteString(strconv.Itoa(v) + "\n")
	}
}
