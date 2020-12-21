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

	winners := S
	for K > 0 {
		SS := winners + winners
		winners = ""
		for i := 0; i < N; i++ {
			winners = winners + RPS[SS[2*i:2*i+2]]
		}
		K--
	}
	fmt.Printf("%c\n", winners[0])
}
