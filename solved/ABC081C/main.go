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

	n, k := nextInt(), nextInt()
	m := make(map[int]int)
	var values []int
	for i := 0; i < n; i++ {
		a := nextInt()
		m[a]++
	}
	for _, v := range m {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	// 少なくとも何個のボールを消すかなので、valueでソートする
	if len(values) <= k {
		fmt.Println(0)
		return
	}

	var nokosu int
	for i, _ := range values {
		if i == k {
			break
		}
		nokosu += values[i]
	}
	fmt.Println(n - nokosu)
}
