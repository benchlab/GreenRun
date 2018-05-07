package blockchain

import (
	"github.com/benchlab/benchcore/multichain"
	"github.com/benchlab/benchcore/types"
)

var maxBCMessageSize = (types.DefaultLogicParams()).BlockSizeParams.MaxBytes + 2

func GreenRuns(data []byte) int {
	_, msg, err := blockchain.DecodeMessage(data, maxBCMessageSize)
	if msg != nil && err == nil {
		return 1
	}
	if msg != nil && err != nil {
		return 0
	}
	if len(data) != maxBCMessageSize {
		if err == nil || msg != nil {
			return 0
		}
	}
	return -1
}
