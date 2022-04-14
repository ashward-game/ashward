package crypto

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CommitmentWithSig struct {
	Commitment string
	Signature  string
}

type Commitment common.Hash

// digest hashes all passed byte slices.
// The passed slices won't be mutated.
func digest(ms ...[]byte) []byte {
	h := crypto.Keccak256Hash(ms...)
	return h.Bytes()
}

// Commit creates a new cryptographic commitment to the passed byte slices
// stuff (which won't be mutated).
func Commit(stuff ...[]byte) Commitment {
	h := digest(stuff...)
	output := common.BytesToHash(h)
	return Commitment(output)
}

// Verify verifies that the underlying commit c was a commit to the passed
// byte slices stuff (which won't be mutated).
func (c *Commitment) Verify(stuff ...[]byte) bool {
	com := common.Hash(*c).Bytes()
	return bytes.Equal(com, digest(stuff...))
}
