package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[int]int)
	var a []int
	for i := 0; i < n; i++ {
		ai := nextInt()
		if m[ai] == 0 {
			a = append(a, ai)
		}
		m[ai]++
	}
	sort.Ints(a)
	//fmt.Println(m)
	//fmt.Println(a)
	ans := 0
	for i := len(a) - 1; i >= 0; i-- {
		if m[a[i]] >= 4 {
			ans = a[i] * a[i]
			break
		}
		if m[a[i]] >= 2 && m[a[i]] < 4 {
			for j := i - 1; j >= 0; j-- {
				if m[a[j]] >= 2 {
					ans = a[i] * a[j]
					break
				}
			}
		}
		if ans > 0 {
			break
		}
	}
	fmt.Println(ans)
}
