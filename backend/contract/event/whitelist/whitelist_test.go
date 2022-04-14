package whitelist

import (
	"os"
	"testing"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var addressFileTemplate = `
{"Marketplace":"{{ .Marketplace }}"}
`
var addressFile = `./address_test.json`

var MkpAddress = common.HexToAddress("0x01")

type addresses struct {
	Marketplace string
}

func generateAddressFile(t *testing.T) func() {
	add := addresses{
		Marketplace: MkpAddress.String(),
	}
	out, err := template.New("addresses").Parse(addressFileTemplate)
	if err != nil {
		t.Fatal(err)
	}
	outFile, err := os.Create(addressFile)
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	err = out.Execute(outFile, add)
	if err != nil {
		t.Fatal(err)
	}
	return func() {
		os.Remove(addressFile)
	}
}

func TestWhiteListMarketplace(t *testing.T) {
	teardown := generateAddressFile(t)
	defer teardown()

	err := Setup(addressFile)
	assert.NoError(t, err)

	isMkp := IsTokenTransferFrom(MkpAddress)
	assert.True(t, isMkp)
	isMkp = IsTokenTransferTo(MkpAddress)
	assert.True(t, isMkp)

	notMkp := IsTokenTransferFrom(common.HexToAddress("0x00"))
	assert.False(t, notMkp)
	notMkp = IsTokenTransferTo(common.HexToAddress("0x00"))
	assert.False(t, notMkp)
}
