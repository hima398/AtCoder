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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Exp(x, y int) int {
	exp := 1
	for i := 0; i < y; i++ {
		exp *= x
	}
	return exp
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Pos struct {
	X int
	Y int
}

var RPS = map[string]string{
	"RR": "R", "RP": "P", "RS": "R",
	"PR": "P", "PP": "P", "PS": "S",
	"SR": "R", "SP": "S", "SS": "S",
}

func Battle(i, j, n int, s string) int {
	a := s[i%n]
	b := s[j%n]

	result := (RPS[a] - RPS[b] + 3) % 3
	if result == 0 || result == 2 {
		//fmt.Printf("Winner %d\n", i)
		return i
	} else {
		// result == 1
		//fmt.Printf("Winner %d\n", j)
		return j
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, K := nextInt(), nextInt()
	S := nextString()

	if len(S) == 1 {
		fmt.Println(S)
		return
	}

	//fmt.Println(winners)
	for len(winners) > 1 {
		var next []int
		for i := 0; i < len(winners)-1; i += 2 {
			winner := Battle(winners[i], winners[i+1], n, s)
			next = append(next, winner)
		}
		winners = next
		//fmt.Println(winners)
	}
	fmt.Printf("%c\n", s[winners[0]%n])
}
