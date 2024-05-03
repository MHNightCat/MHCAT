package pkg

import "math/big"

var (
	// Red
	ErrorColor = HexToInt("#FF4242")
	// Green
	SuccessfulColor = HexToInt("#66FF25")
)

func HexToInt(color string) int {
	n := new(big.Int)
	n.SetString(color[1:], 16)
	return int(n.Int64())
}
