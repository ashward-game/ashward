package crypto

import (
	"crypto/ecdsa"
	"orbit_nft/util"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
)

// Input: Zero or many slice of byte
// Ouput: Slice of byte after concatnated all the slice passed to the function,
//        with the ethereum prefix `"\x19Ethereum Signed Message:\n" + message.length`
// Note: Passing parameters ([]byte("ABC"),[]byte("XYZ"))
//       will produce the same hash as passing ([]byte("ABCX"),[]byte("YZ")) since parameters are concatnated
func ethPrefix(msgs ...[]byte) (rs []byte) {
	var length int
	for _, slice := range msgs {
		rs = append(rs, slice...)
		length += len(slice)
	}
	rs = append([]byte("\x19Ethereum Signed Message:\n"+strconv.Itoa(length)), rs...)
	return
}

func EthSign(prv *ecdsa.PrivateKey, msg []byte) (sig []byte, err error) {
	sig, err = crypto.Sign(HashMessage(msg), prv)
	// see https://ethereum.stackexchange.com/questions/102190/signature-signed-by-go-code-but-it-cant-verify-on-solidity
	if sig[64] == 0 || sig[64] == 1 {
		sig[64] += 27
	}
	return
}

func EthSigVerify(pk, msg, sig []byte) bool {
	sigNoRecoverID := sig[:len(sig)-1]
	return crypto.VerifySignature(pk, HashMessage(msg), sigNoRecoverID)
}

func SaveAddress(file string, key ecdsa.PublicKey, accountName string) error {
	k := crypto.PubkeyToAddress(key).Hex()
	return util.WriteToJSONFile(file, accountName, k)
}

func SaveMnemonic(file string, mnemonic string) error {
	return util.WriteToJSONFile(file, "mnemonic", mnemonic)
}

func GetPrivateKey(file string) (*ecdsa.PrivateKey, error) {
	return crypto.LoadECDSA(file)
}

func HashMessage(msg []byte) []byte {
	return crypto.Keccak256(ethPrefix(msg))
}
