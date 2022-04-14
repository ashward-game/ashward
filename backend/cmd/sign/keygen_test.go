package sign

import (
	"crypto/ecdsa"
	"io/ioutil"
	cryp "orbit_nft/crypto"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestVerifySignature(t *testing.T) {
	// generate sign key
	pathSecrets := "./"
	addressFile := filepath.Join(pathSecrets, "address.json")
	filePrivateKey := filepath.Join(pathSecrets, "sign.private")
	walletFile := filepath.Join(pathSecrets, "wallet.json")
	defer func() {
		os.Remove(filePrivateKey)
		os.Remove(addressFile)
		os.Remove(walletFile)
	}()

	ioutil.WriteFile(addressFile, []byte("{}"), 0600)
	err := generate(pathSecrets, addressFile)
	assert.NoError(t, err)

	privateKey, err := crypto.LoadECDSA(filePrivateKey)
	assert.NoError(t, err)

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	assert.True(t, ok)
	PublicKey := crypto.FromECDSAPub(publicKeyECDSA)
	msg := []byte("sign this message")

	//signing
	eth_sign, err := cryp.EthSign(privateKey, msg)
	assert.NoError(t, err)

	// verify
	success := cryp.EthSigVerify(PublicKey, msg, eth_sign)
	assert.True(t, success)
}
