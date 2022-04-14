package sign

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"orbit_nft/constant"
	oCrypto "orbit_nft/crypto"

	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
)

// KeygenCmd represents the keygen command
var KeygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generate a keypair for signing",
	Long:  `Generate a keypair for signing and store the generated key pair to files.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("out")
		if err != nil {
			log.Fatal(err)
		}

		addressFile, err := cmd.Flags().GetString("addressFile")
		if err != nil {
			log.Fatal(err)
		}

		err = generate(path, addressFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	KeygenCmd.Flags().StringP("out", "o", "./secrets/", "Path to directory storing the generated key pair")
	KeygenCmd.Flags().String("addressFile", "../common/address.json", "Path to the file containing all common addresses")
}

func generate(path, addressFile string) error {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return err
	}

	derivationPath := hdwallet.MustParseDerivationPath(constant.EthDerivationPath)
	account, err := wallet.Derive(derivationPath, true)
	if err != nil {
		return err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return err
	}

	publicKey := privateKey.PublicKey

	// store private/public-key to files
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	privateKeyPath := filepath.Join(path, "sign.private")
	walletPath := filepath.Join(path, "wallet.json")

	if _, err := os.Stat(privateKeyPath); err == nil {
		return errors.New("private key file is already exists. Exiting")
	}

	// create wallet.json if not exist
	if _, err := os.Stat(privateKeyPath); errors.Is(err, os.ErrNotExist) {
		if err := ioutil.WriteFile(walletPath, []byte("{}"), 0600); err != nil {
			return err
		}
	}

	if err := crypto.SaveECDSA(privateKeyPath, privateKey); err != nil {
		return err
	}
	// save to address.json using Backend as the key
	if err := oCrypto.SaveAddress(addressFile, publicKey, "Backend"); err != nil {
		return err
	}
	if err := oCrypto.SaveMnemonic(walletPath, mnemonic); err != nil {
		return err
	}
	return nil
}
