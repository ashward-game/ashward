package crypto

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common"
)

// MakeRand returns a random slice of bytes.
// It returns an error if there was a problem while generating
// the random slice.
// It is different from the 'standard' random byte generation as it
// hashes its output before returning it; by hashing the system's
// PRNG output before it is send over the wire, we aim to make the
// random output less predictable (even if the system's PRNG isn't
// as unpredictable as desired).
// See https://trac.torproject.org/projects/tor/ticket/17694
func MakeRand() ([]byte, error) {
	r := make([]byte, common.HashLength)
	if _, err := rand.Read(r); err != nil {
		return nil, err
	}
	// Do not directly reveal bytes from rand.Read on the wire
	return digest(r), nil
}
