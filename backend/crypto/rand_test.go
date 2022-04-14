package crypto

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type testErrorRandReader struct{}

func (er testErrorRandReader) Read([]byte) (int, error) {
	return 0, errors.New("Not enough entropy!")
}

func TestMakeRand(t *testing.T) {
	r, err := MakeRand()
	if err != nil {
		t.Fatal(err)
	}
	// check if hashed the random output:
	if len(r) != common.HashLength {
		t.Fatal("Looks like Digest wasn't called correctly.")
	}
	orig := rand.Reader
	rand.Reader = testErrorRandReader{}
	_, err = MakeRand()
	if err == nil {
		t.Fatal("No error returned")
	}
	rand.Reader = orig
}
