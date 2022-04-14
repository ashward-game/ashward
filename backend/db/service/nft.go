package service

import (
	"errors"
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
	"orbit_nft/db/util"
)

type TokenFilter struct {
	Type         string
	Class        string
	Rarity       string
	Search       string
	OrderByPrice string
}

type NftService struct {
	repo *repository.NftRepository
}

func NewNftService(repo *repository.NftRepository) *NftService {
	return &NftService{
		repo: repo,
	}
}

func (s *NftService) GetOnSaleTokensWithFilter(pg *util.Pagination, filter TokenFilter) (uint, []model.NFT, error) {
	// sanity check
	var err error

	query := s.repo.FindTokens(pg).
		OnSale().
		WithType(filter.Type).
		WithRarity(filter.Rarity).
		WithClass(filter.Class)

	type record struct {
		model.NFT
		Total uint
	}

	var results []record

	err = query.ByNameOrTokenId(filter.Search).
		SortByPrice(filter.OrderByPrice).
		Find(&results)
	if err != nil {
		return 0, nil, err
	}
	if len(results) == 0 {
		return 0, make([]model.NFT, 0), nil
	}
	var actualResults []model.NFT
	for _, v := range results {
		actualResults = append(actualResults, v.NFT)
	}

	return results[0].Total, actualResults, nil
}

func (s *NftService) GetTokenById(tokenId uint) (*model.NFT, error) {
	nft, err := s.repo.FindByTokenID(tokenId).WithTradingHistory().FirstToken()
	if err != nil {
		return nil, err
	}

	return nft, nil
}

func (s *NftService) CreateNotForSaleToken(tokenId uint, owner string, metadataURI string, rawStringMetadata string) error {
	nft := model.NewNFToken(tokenId, owner, metadataURI, rawStringMetadata)

	return s.repo.CreateNotForSaleToken(nft)
}

func (s *NftService) TransferNFT(tokenId uint, from string, to string) error {
	nft, err := s.repo.FindByTokenID(tokenId).FirstToken()
	if err != nil {
		return err
	}

	if nft.IsForSale() {
		return errors.New(ErrNftIsOnMarketplace)
	}

	// cheat: if `from` is empty, we intentionally ignore this sanity check, as it could be the case that a related event was ignored.
	if len(from) > 0 && nft.Owner != from {
		return errors.New(ErrNftTransferFromNotFound)
	}
	nft.Owner = to

	return s.repo.Update(nft)
}

func (s *NftService) GetTokensOfAddressWithFilter(pg *util.Pagination, address string, filter TokenFilter) (uint, []model.NFT, error) {
	// sanity check
	var err error

	query := s.repo.FindTokens(pg).
		OwnedBy(address).
		WithType(filter.Type).
		WithRarity(filter.Rarity).
		WithClass(filter.Class)

	type record struct {
		model.NFT
		Total uint
	}

	var results []record

	err = query.ByNameOrTokenId(filter.Search).
		SortByPrice(filter.OrderByPrice).
		Find(&results)
	if err != nil {
		return 0, nil, err
	}
	if len(results) == 0 {
		return 0, make([]model.NFT, 0), nil
	}
	var actualResults []model.NFT
	for _, v := range results {
		actualResults = append(actualResults, v.NFT)
	}

	return results[0].Total, actualResults, nil
}
