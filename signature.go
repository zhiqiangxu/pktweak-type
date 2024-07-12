package pktweaktype

import "math/big"

type SigFlavor int

const (
	ETH SigFlavor = iota
	BTC
	STD
	SUPR
)

type Signature struct {
	Flavor SigFlavor
	Value  interface{}
}

type ECDSAEthSig []byte

type ECDSAStdSig struct {
	R, S *big.Int
}

// TODO: convert between signature flavors
