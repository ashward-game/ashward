package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestEthPrefix(t *testing.T) {
	// test for 32 bytes data input

	//message to hash
	msg := "test message"
	// the hash will have 32 bytes Length
	hash := HashMessage([]byte(msg))

	// use the hash value for the test
	expect := append([]byte("\x19Ethereum Signed Message:\n32"), hash...)
	ethhash := ethPrefix(hash)

	//assert equal
	if !bytes.Equal(expect, ethhash) {
		t.Errorf("EthPrefix error: \n expect data: %q \n but got data: %q", expect, ethhash)
	}

	// test for empty data
	ethhash = ethPrefix()
	expect = []byte("\x19Ethereum Signed Message:\n0")
	if !bytes.Equal(ethhash, expect) {
		t.Fatalf("EthPrefix error: \n expect data: %q \n but got data: %q", expect, ethhash)
	}
}

func TestEthSignVerify(t *testing.T) {
	//prepare keys and data for the test
	prv, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	publicKeyECDSA, ok := prv.Public().(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("error casting public key to ECDSA")
	}
	PublicKey := crypto.FromECDSAPub(publicKeyECDSA)
	msg := []byte("sign this message")

	//signing
	eth_sign, err := EthSign(prv, msg)
	if err != nil {
		t.Fatal(err)
	}

	// verify
	if !EthSigVerify(PublicKey, msg, eth_sign) {
		t.Fatal("EthVerify Fail to verify")
	}
}

func TestEcrecover(t *testing.T) {
	//prepare keys and data for the test
	prv, err := crypto.GenerateKey()
	assert.NoError(t, err)

	publicKeyECDSA, ok := prv.Public().(*ecdsa.PublicKey)
	assert.True(t, ok)

	PublicKey := crypto.FromECDSAPub(publicKeyECDSA)
	msg := []byte("sign this message")

	//signing
	eth_sign, err := EthSign(prv, msg)
	assert.NoError(t, err)

	// verify
	assert.True(t, EthSigVerify(PublicKey, msg, eth_sign))

	// Transform yellow paper V from 27/28 to 0/1
	// see https://ethereum.stackexchange.com/questions/102190/signature-signed-by-go-code-but-it-cant-verify-on-solidity
	// and https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
	eth_sign[64] -= 27

	//ecrecover
	pb, err := crypto.Ecrecover(HashMessage(msg), eth_sign)
	assert.NoError(t, err)

	assert.True(t, bytes.Equal(PublicKey, pb))
}
