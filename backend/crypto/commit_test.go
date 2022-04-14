package crypto

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestDigest(t *testing.T) {
	msg := []byte("test message")
	d := digest(msg)
	if len(d) != common.HashLength {
		t.Fatal("Computation of Hash failed.")
	}
	if bytes.Equal(d, make([]byte, common.HashLength)) {
		t.Fatal("Hash is all zeros.")
	}
}

func TestCommit(t *testing.T) {
	stuff := []byte("123")
	commit := Commit(stuff)
	if !commit.Verify(stuff) {
		t.Fatal("Commit with salt doesn't verify!")
	}
}
