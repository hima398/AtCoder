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
	x := make([]int, n)
	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}
	copy(sorted, x)
	sort.Ints(sorted)
	m1, m2 := sorted[n/2-1], sorted[n/2]
	// m1 == m2であればm1かm2を出力する処理を追加で書いても良いけどもあまり時間は変わらないかもしれない？
	for _, v := range x {
		if v <= m1 {
			fmt.Println(m2)
		} else {
			// v > m1
			fmt.Println(m1)
		}
	}
}
