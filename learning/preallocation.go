package learning

import (
	"fmt"
	"time"
)

func Preallo() {
	var n = 100000000
	var testSlice = []int{}
	var testSlicePreAllo = make([]int, 0, n)
	fmt.Printf("testsclice = %v\n", TestIt(testSlice, n))
	fmt.Printf("testsclice = %v\n", TestIt(testSlicePreAllo, n))
}

func TestIt(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
