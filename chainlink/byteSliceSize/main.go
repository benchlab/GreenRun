package bsz

import (
	"fmt"

	"github.com/benchlab/chainlink"
)

func GreenRuns(data []byte) int {
	n := chainlink.ByteSliceSize(data)
	if n < 1 {
		panic("expecting at least 1")
	}
	switch {
	case n == 1 && len(data) != 0:
		panic(fmt.Sprintf("n==0 yet len(data)==%d", len(data)))
	case n != 1 && len(data) == 0:
		panic(fmt.Sprintf("n==%d yet len(data)==0", n))
	default:
		return 1
	}
}
