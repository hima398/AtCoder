package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s); j++ {
			ss := s[:i] + s[j:]
			if ss == "keyence" {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
