package main

import (
	"flag"
	"fmt"
)

func fiboSeries(n int) (fibos []int) {
	fibos = append(fibos, 1)
	if n == 0 {
		return
	}
	if n > 0 {
		fibos = append(fibos, 1)

	}
	if n == 1 {
		return
	}
	for i := 2; i <= n; i++ {
		fibos = append(fibos, fibos[i-1]+fibos[i-2])
	}
	return
}

func main() {
	fiboNum := flag.Int("num", 0, "Number of fibonacci numbers to print")
	flag.Parse()
	fmt.Printf("fibo series for first %v ints is %v ", *fiboNum+1, fiboSeries(*fiboNum))
}
