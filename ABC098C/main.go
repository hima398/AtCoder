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

func CumulativeSum(n int, row string) int {
	// 番兵を持たせる
	w := make([]int, n+2)
	e := make([]int, n+2)
	for i := 0; i < n; i++ {
		if row[i] == 'W' {
			w[i+1] = w[i] + 1
			e[i+1] = e[i]
		} else if row[i] == 'E' {
			w[i+1] = w[i]
			e[i+1] = e[i] + 1
		} else {
			w[i+1] = w[i]
			e[i+1] = e[i]
		}
	}
	result := n
	for i := 0; i < n; i++ {
		count := w[i] + (e[n] - e[i])
		if result > count {
			result = count
		}
	}
	return result
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	row := nextString()

	fmt.Println(CumulativeSum(n, row))

}
