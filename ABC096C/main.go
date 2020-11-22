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

	// 1 <= h, w <= 50
	h, w := nextInt(), nextInt()

	board := make([]string, h)
	for i := 0; i < h; i++ {
		board[i] = nextString()
	}

	kx := []int{1, 0, -1, 0}
	ky := []int{0, 1, 0, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == '.' {
				continue
			}
			canDraw := false
			for k := 0; k < 4; k++ {
				ni := i + ky[k]
				nj := j + kx[k]
				if ni < 0 || ni >= h {
					continue
				}
				if nj < 0 || nj >= w {
					continue
				}
				if board[ni][nj] == '#' {
					canDraw = true
					break
				}
			}
			if canDraw == false {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
