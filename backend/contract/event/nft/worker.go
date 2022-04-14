package nft

import (
	"context"
	"io"
	"math/big"
	"net/http"
	"orbit_nft/contract"
	"orbit_nft/contract/abi/nft"
	orbitContext "orbit_nft/contract/context"
	serviceNft "orbit_nft/contract/service/nft"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

var _getTokenURI = func(ctx context.Context, tokenId *big.Int) (string, error) {
	addressFile := ctx.Value(orbitContext.KeyAddressFile).(string)
	chainId := ctx.Value(orbitContext.KeyChainId).(string)
	cli, err := contract.NewBscClient(addressFile, chainId)
	if err != nil {
		return "", err
	}
	defer cli.Close()
	addressNFT, err := util.GetContractAddress(cli.AddressFile(), nft.Name)
	if err != nil {
		return "", err
	}

	addressContract := common.HexToAddress(addressNFT)
	nftContract, err := serviceNft.NewNFToken(addressContract, cli.Client())
	if err != nil {
		return "", err
	}
	return nftContract.TokenURI(tokenId)
}

func getTokenURI(ctx context.Context, tokenId *big.Int) (string, error) {
	return _getTokenURI(ctx, tokenId)
}

var _getTokenMetadata = func(ctx context.Context, tokenUri string) (string, error) {
	resp, err := http.Get(tokenUri)
	if err != nil {
		return "", err
	}
	defer func() {
		resp.Header.Set("Connection", "close")
		resp.Close = true
		resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getTokenMetadata(ctx context.Context, tokenUri string) (string, error) {
	return _getTokenMetadata(ctx, tokenUri)
}
