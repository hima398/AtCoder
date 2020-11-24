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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[int]int)
	var result int
	for i := 1; i <= n; i++ {
		a := nextInt()
		// 紙に書かれていたら削除
		if m[a] != 0 {
			m[a] = 0
			result--
		} else {
			m[a] = i
			result++
		}
	}
	fmt.Println(result)
}
