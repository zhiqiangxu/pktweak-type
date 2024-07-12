package pktweaktype

import "math/big"

type Tweaker interface {
	// converts real pk to tweaked pk
	Tweak(realPK, tweak *big.Int) *big.Int

	// initialize by tweaked pk and tweak pk
	Initialize(tweakedPK, tweak *big.Int) error
	// sign a flavored signature like real pk,
	// the exact flavor is opinionatedly determinted by curve, users can convert signature between flavors though.
	// Note: the caller is responsible for choosing a proper hash function and compute the hash digest
	Sign(hash []byte) (*Signature, error)
}
