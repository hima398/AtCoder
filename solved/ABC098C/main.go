package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Honestly(n int, row string) int {
	result := n
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			if j < i && row[j] == 'W' {
				count++
			}
			if j > i && row[j] == 'E' {
				count++
			}
		}
		if result > count {
			result = count
		}
	}
	return result
}

func Solve(n int, s string) int {
	se := make([]int, n+2) // Eの累積和
	sw := make([]int, n+2) // Wの累積和

	for i := 1; i <= n; i++ {
		if s[i-1] == 'E' {
			se[i] = se[i-1] + 1
			sw[i] = sw[i-1]
		} else if s[i-1] == 'W' {
			se[i] = se[i-1]
			sw[i] = sw[i-1] + 1
		}
	}
	se[n+1] = se[n]
	//fmt.Println(se)
	//fmt.Println(sw)
	ans := n
	for i := 1; i <= n; i++ {
		//リーダーの左側にいるWの数
		sum := sw[i-1]
		//リーダーの右側にいるEの数
		sum += se[n+1] - se[i]
		ans = Min(ans, sum)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	row := nextString()

	//fmt.Println(Honestly(n, row))
	fmt.Println(Solve(n, row))
}
