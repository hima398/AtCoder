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

var dp [2000][2000]int
var x, y, z [2000][2000]int

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	board := make([]string, h)
	for i := 0; i < h; i++ {
		board[i] = nextString()
	}

	kx := []int{1, 0, -1, 0, 1, -1, -1, 1}
	ky := []int{0, 1, 0, -1, 1, 1, -1, -1}
	for i := 0; i < h; i++ {
		row := ""
		for j := 0; j < w; j++ {
			if board[i][j] == '#' {
				row += "#"
				continue
			}
			num := 0
			for k := 0; k < 8; k++ {
				ni := i + ky[k]
				nj := j + kx[k]
				if ni < 0 || ni >= h {
					continue
				}
				if nj < 0 || nj >= w {
					continue
				}
				if board[ni][nj] == '#' {
					num++
				}
			}
			row += strconv.Itoa(num)
		}
		board[i] = row
	}

	for i := 0; i < h; i++ {
		fmt.Println(board[i])
	}
}
