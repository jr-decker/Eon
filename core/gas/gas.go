package gas

import "math/big"

// Gas (measured in Virtual Gas Units) are the fees used on the layer1 blockchain.
// 1 Vgu = 10^7 Lunch.
type Gas uint64

type GasPrice *big.Int
