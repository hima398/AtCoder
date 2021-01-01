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

// 公式解説の写経
func SolveCommentary(n, k int, s string) int {
	// 1がx0回、0がx1回、、、というダイジェストを作る
	now := byte(1)
	cnt := 0
	var nums []int
	for i := 0; i < n; i++ {
		if s[i] == byte('0')+now {
			cnt++
		} else {
			nums = append(nums, cnt)
			now ^= 1
			cnt = 1
		}
	}
	if cnt != 0 {
		nums = append(nums, cnt)
	}
	if len(nums)%2 == 0 {
		nums = append(nums, 0)
	}

	add := 2*k + 1
	l := 0
	r := 0
	sum := 0 // [l, r)のsum
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		NextL := i
		NextR := Min(i+add, len(nums))

		for NextL > l {
			sum -= nums[l]
			l++
		}
		for NextR > r {
			sum += nums[r]
			r++
		}
		ans = Max(ans, sum)
	}
	return ans
}

/*
func Solve(n, k int, s string) int {
	i := 0
	// 最初の0までインクリメント
	for i < n && s[i] == '1' {
		i++
	}

		j := i
		for p := 0; p < k; p++ {
			for j < n && s[j] == '0' {
				j++
			}
			for j < n && s[j] == '1' {
				j++
			}
		}
	ans := j
	b := 0
	for j < n {
		for b < n && s[b] == '1' {
			b++
		}
		for b < n && s[b] == '0' {
			b++
		}
		for j < n && s[j] == '0' {
			j++
		}
		for j < n && s[j] == '1' {
			j++
		}
		tmp := j - b
		ans = Max(ans, tmp)
	}
	return ans
}
*/

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	s := nextString()

	fmt.Println(SolveCommentary(n, k, s))

}
