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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a := nextInt()
		if a%4 == 0 {
			m[4]++
		} else if a%2 == 0 {
			m[2]++
		} else {
			m[1]++
		}
	}
	// m[2] == 0
	if m[2] == 0 && m[1] <= m[4]+1 {
		fmt.Println("Yes")
		return
	}
	if m[1] <= m[4] {
		fmt.Println("Yes")
		return
	} else {
		fmt.Println("No")
		return
	}
}
